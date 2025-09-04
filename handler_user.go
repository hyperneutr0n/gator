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
		return errors.New("user " + cmd.Args[0] + " not found" + "\n" + err.Error())
	}
	
	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
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
		return errors.New("failed to register user " + cmd.Args[0] + "\n" + err.Error())
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return errors.New("failed to set user to config file" + "\n" + err.Error())
	}

	fmt.Println("User " + user.Name + " has been created")
	return nil
}