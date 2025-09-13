package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hyperneutr0n/rss-aggregator/internal/database"
)

var ErrCreatePost = errors.New("failed creating post")

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("agg expects time duration string. example: 1s, 1m, 1h")
	}
	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed parsing parse duration: %w", err)
	}

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			if !errors.Is(err, ErrCreatePost) {
				return fmt.Errorf("%w", err)
			} else {
				fmt.Println(err)
			}
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeeds, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed fetching next feed: %w", err)
	}

	layout := time.RFC1123Z
	for _, nextFeed := range nextFeeds {
		err := s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
		if err != nil {
			return fmt.Errorf("failed marking feed as fetched: %w", err)
		}

		rss, err := fetchFeed(context.Background(), nextFeed.Url)
		if err != nil {
			return fmt.Errorf("failed to fetch feed: %w", err)
		}

		fmt.Println(nextFeed.Name)
		for _, item := range rss.Channel.Item {
			pubAt, err := time.Parse(layout, item.PubDate)
			if err != nil {
				return fmt.Errorf("failed parsing time: %w", err)
			}

			post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
				FeedID: nextFeed.ID,
				Title: item.Title,
				Url: item.Link,
				Description: item.Description,
				PublishedAt: pubAt,
			})
			if err != nil {
				return fmt.Errorf("%w: %w", ErrCreatePost, err)
			}
			fmt.Println(post.Title)
		}
	}
	return nil
}
