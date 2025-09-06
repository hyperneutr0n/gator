package main

import (
	"context"
	"fmt"

	"github.com/hyperneutr0n/rss-aggregator/internal/database"
)

func isLoggedIn(handler func(s *state, cmd command, user database.User) error) (func(*state, command) error) {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return fmt.Errorf("error when looking up your id: %w", err)
		}
		return handler(s, cmd, user)
	}
}