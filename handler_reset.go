package main

import (
	"fmt"
	"context"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.RemoveUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not delete users: %w", err)
	}
	fmt.Println("database reset successful")
	return nil
}