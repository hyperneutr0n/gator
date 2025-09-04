package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 1 {
		return errors.New("login expects a single argument, the username")
	}

	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("user " + cmd.Args[0] + " not found: %w", err)
	}
	
	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("failed setting user to config file: %w", err)
	}
	fmt.Println("User " + user.Name + " been set. Login success!")
	
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("register expects name as an argument")
	}
	
	user, err := s.db.CreateUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("failed to set user to config file: %w", err)
	}

	fmt.Println("User " + user.Name + " has been created")
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.ResetUser(context.Background()); err != nil {
		return fmt.Errorf("failed resetting user table: %w", err)
	}
	
	fmt.Println("Successfully resetting user table")
	return nil
}