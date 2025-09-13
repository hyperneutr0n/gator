package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperneutr0n/rss-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.Args[0]) > 0 {
		tmp, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("failed parsing limit: %w", err)
		}
		limit = int32(tmp)
	}

	posts, err := s.db.GetPostFromUser(context.Background(), database.GetPostFromUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("failed fetching posts from database: %w", err)
	}

	for _, post := range posts {
		printPost(post)
		printLineLimitter()
	}
	return nil
}

func printPost(post database.Post) {
	layout := time.RFC850
	fmt.Printf("%v - %v\n", post.Title, post.ID)
	fmt.Printf("%v\n", post.PublishedAt.Format(layout))
	fmt.Printf("%v\n", post.Url)
	fmt.Printf("%v\n", post.Description)
}