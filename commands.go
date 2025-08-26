package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	callbacks map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.callbacks[cmd.Name]
	if !ok {
		return errors.New("Command " + cmd.Name + " not found")
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.callbacks[name] = f
}
