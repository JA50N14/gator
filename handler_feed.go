package main

import (
	"fmt"
	"context"
	"time"

	"github.com/JA50N14/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	createFeedParams := database.CreateFeedParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
		Url: cmd.Args[1],
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), createFeedParams)
	if err != nil {
		return fmt.Errorf("could not post feed in database: %w", err)
	}

	createFeedFollowParams := database.CreateFeedFollowParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		return fmt.Errorf("could not insert into feed follows: %w", err)
	}
	printFeedRecord(feed)
	return nil
}


func printFeedRecord(f database.Feed) {
	fmt.Printf("ID: %v\nCreatedAt: %v\nUpdatedAt: %v\nName: %s\nUrl: %s\nUserID: %v\n=================================\n", f.ID, f.CreatedAt, f.UpdatedAt, f.Name, f.Url, f.UserID)
}


func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not get list of feeds: %w", err)
	}
	if len(feeds) == 0 {
		fmt.Println("No feeds in database")
		return nil
	}
	
	for _, feed := range feeds {
		fmt.Println("Feed Name:", feed.Feedname)
		fmt.Println("URL:", feed.Url)
		fmt.Println("User:", feed.Username)
		fmt.Println("=====================")
	}
	return nil
}




