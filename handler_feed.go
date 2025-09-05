package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hyperneutr0n/rss-aggregator/internal/database"
)


func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return errors.New("addFeed expects name and url as an arguments")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("failed fetching current user's id: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		UserID: user.ID,
		Name: cmd.Args[0],
		Url: cmd.Args[1],
	})
	if err != nil {
		return fmt.Errorf("failed creating feed record: %w", err)
	}

	printFeed(feed)

	if err := followFeed(s, user.ID, feed.ID); err != nil {
		return err
	}
	
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed fetching feeds from database: %w", err)
	}
	
	for _, feed := range feeds {
		printFeed(database.Feed{
			ID: feed.ID,
			UserID: feed.UserID,
			Name: feed.Name,
			Url: feed.Url,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		})
		fmt.Printf("*User Name:		%v\n", feed.UserName)
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:			%v\n", feed.ID)
	fmt.Printf("* Name:			%v\n", feed.Name)
	fmt.Printf("* URL:			%v\n", feed.Url)
	fmt.Printf("* User ID:		%v\n", feed.UserID)
}

func handlerFollow (s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("follow command expect a feed's link as an argument")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error when looking up existing feed: %w", err)
	}
	
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error when looking up current user's id: %w", err)
	}

	if err := followFeed(s, user.ID, feed.ID); err != nil {
		return err
	}
	
	return nil
}

func handlerFollowing (s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error when looking up your id: %w", err)
	}

	followedFeeds, err := s.db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error when fetching your followed feeds: %w", err)
	}

	fmt.Println("These are feeds that you followed:")
	for _, feed := range followedFeeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}

func followFeed(s *state, userID uuid.UUID, feedID int32) error {
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: userID,
		FeedID: feedID,
	})
	if err !=nil {
		return fmt.Errorf("error when creating feed follows record: %w", err)
	}

	fmt.Println("Successfully following " + feedFollow.FeedName)
	
	return nil
}