package main

import (
	"fmt"
	"context"

	"github.com/JA50N14/gator/internal/database"
)


func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	deleteFeedFollowForUserParams := database.DeleteFeedFollowForUserParams {
		Url: cmd.Args[0],
		UserID: user.ID,
	}

	err := s.db.DeleteFeedFollowForUser(context.Background(), deleteFeedFollowForUserParams)
	if err != nil {
		return fmt.Errorf("unable to delete feed: %w", err)
	}
	fmt.Printf("No longer following feed - %s\n", cmd.Args[0])
	return nil
}