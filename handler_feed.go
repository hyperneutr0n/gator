package main

import (
	"context"
	"errors"
	"fmt"

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

	fmt.Println(feed)
	return nil
}