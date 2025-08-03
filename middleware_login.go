package main

import (
	"fmt"
	"context"

	"github.com/JA50N14/gator/internal/database"
)


func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user not logged in or registered: %w", err)
		}
		return handler(s, cmd, user)
	}
}