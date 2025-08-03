package main

import (
	"fmt"
	"strings"
	"time"
	"os"
	"context"
	"database/sql"

	"github.com/JA50N14/gator/internal/database"
	"github.com/lib/pq"
	"github.com/google/uuid"
)


func handlerRegister (s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: register <name>")
	}
	createUserParams := database.CreateUserParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	}
	user, err := s.db.CreateUser(context.Background(), createUserParams)
	
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				os.Exit(1)
			}
		}
		return err
	}

	s.cfg.SetUser(user.Name)
	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}


func handlerLogin (s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		if err == sql.ErrNoRows {
			os.Exit(1)
		}
		return err
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}
	fmt.Printf("User: %s has been set\n", strings.Join(cmd.Args, ""))
	return nil
}


func handlerGetUsers(s *state, cmd command) error {
	var nameList []string
	nameList, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not list users: %w", err)
	}
	for _, name := range nameList {
		if name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", name)
		} else {
			fmt.Printf("* %s\n", name)
		}
	}
	return nil
}


func printUser(user database.User) {
	fmt.Printf("Name: %s\nID: %v\nCreatedAt: %v\nUpdatedAt: %v\n", user.Name, user.ID, user.CreatedAt, user.UpdatedAt)
}