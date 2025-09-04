package main

import (
	"context"
	"fmt"
)


func handlerAgg(s *state, cmd command) error {
	// if len(cmd.Args) < 1 {
	// 	return errors.New("agg expects feed url as an argument")
	// }

	rssFeeds, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed fetching feeds: %w", err)
	}

	fmt.Println(rssFeeds)
	return nil
}
