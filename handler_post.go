package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

	shellCmd := exec.Command("stty", "size")
	shellCmd.Stdin = os.Stdin
	out, err := shellCmd.Output()
	if err != nil {
		fmt.Println("failed getting terminal width")
	}

	terminalSize := strings.Split(strings.TrimSpace(string(out)), " ")
	width, err := strconv.Atoi(terminalSize[1])
	if err != nil {
		fmt.Println("failed parsing width")
	}
	printLineLimitter(width)

	for _, post := range posts {
		printPost(post)
		printLineLimitter(width)
	}
	return nil
}

func printLineLimitter(width int) {
	length := 30
	if width != 0 {
		length = width
	}
	line := ""
	for range length {
		line += "="
	}
	fmt.Println(line)
}

func printPost(post database.Post) {
	layout := time.RFC850
	fmt.Printf("%v - %v\n", post.Title, post.ID)
	fmt.Printf("%v\n", post.PublishedAt.Format(layout))
	fmt.Printf("%v\n", post.Url)
	fmt.Printf("%v\n", post.Description)
}