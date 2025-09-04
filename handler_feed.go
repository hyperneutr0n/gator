package main

import (
	"context"
	"errors"
	"fmt"
	"time"

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