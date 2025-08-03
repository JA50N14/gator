package main

import (
	"fmt"
	"context"
	"strconv"

	"github.com/JA50N14/gator/internal/database"
)


func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil || len(cmd.Args) > 1 {
			return fmt.Errorf("usage: %s <int>", cmd.Name)
		}
		limit = parsedLimit
	}
	
	getPostsForUserParams := database.GetPostsForUserParams {
		UserID: user.ID,
		Limit: int32(limit),
	}

	userPosts, err := s.db.GetPostsForUser(context.Background(), getPostsForUserParams)
	if err != nil {
		return fmt.Errorf("issue calling GetPostsForUser: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(userPosts), user.Name)
	for _, post := range userPosts {
		if post.Title != "" {
			fmt.Printf("*Title: %s\n", post.Title)
		} else {
			fmt.Println("*Title: (no title)")
		}
		
		if post.PublishedAt.Valid {
			fmt.Printf("*Published at: %v\n", post.PublishedAt.Time.Format("2006/01/02"))
		} else {
			fmt.Println("Published at: (no date)")
		}

		fmt.Println("-----------------------------------")
	}
	return nil
}