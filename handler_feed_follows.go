package main

import (
	"fmt"
	"time"
	"context"

	"github.com/JA50N14/gator/internal/database"
	"github.com/google/uuid"
)


func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Errorf("cannot follow feed because it does not exist: %w", err)
	}

	createFeedFollowParams := database.CreateFeedFollowParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		return fmt.Errorf("could not insert into feed_follows table: %w", err)
	}
	printFeedFollowRow(feedFollowRow)
	return nil
}


func printFeedFollowRow(ff database.CreateFeedFollowRow) {
	fmt.Println("Feed Name:", ff.Feedname)
	fmt.Println("Current User:", ff.Username)
}


func handlerListFeedFollows (s *state, cmd command, user database.User) error {
	feedFollowsForUserRows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get feed follows for user: %w", err)
	}
	for _, feedFollowRow := range feedFollowsForUserRows {
		fmt.Println(feedFollowRow.Feedname)
	}
	return nil
}