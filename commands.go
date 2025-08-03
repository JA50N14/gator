package main

import (
	"errors"
)


type command struct {
	Name string
	Args []string
}


type commands struct {
	cmdMap map[string]func(*state, command) error
}


func (c *commands) run(s *state, cmd command) error {
	handlerFunc, ok := c.cmdMap[cmd.Name]
	if !ok {
		return errors.New("command does not exist")
	}

	return handlerFunc(s, cmd)
}


func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdMap[name] = f
}


