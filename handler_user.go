package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 1 {
		return errors.New("Login expects a single argument, the username")
	}

	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Println("User " + cmd.Args[0] + " been set. Login success!")
	
	return nil
}