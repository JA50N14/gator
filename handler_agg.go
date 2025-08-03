package main

import (
	"fmt"
	"context"
	"time"
	"log"
	"database/sql"
	
	"github.com/lib/pq"
	"github.com/JA50N14/gator/internal/database"
	"github.com/google/uuid"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration (i.e. 1s, 2s, etc.): %w", err)
	}
	log.Printf("Collecting feeds every %v\n", time_between_reqs)
	
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	return nil
}


func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("could not get next feeds to fetch", err)
		return
	}

	log.Println("Found a feed to fetch!")

	lastFetched := sql.NullTime {
		Time: time.Now(),
		Valid: true,
	}
	
	markFeedFetchedParams := database.MarkFeedFetchedParams {
		LastFetchedAt: lastFetched,
		UpdatedAt: time.Now(),
		ID: feed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), markFeedFetchedParams)
	if err != nil {
		log.Printf("could not mark %s as being fetched: %v", feed.Name, err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("could not collect feed %s: %v", feed.Name, err)
		return
	}
	
	for _, item := range rssFeed.Channel.Item {
		var postPubDate sql.NullTime
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil || t.IsZero() {
			postPubDate.Valid = false
		} else {
			postPubDate.Time = t
			postPubDate.Valid = true
		}

		var postDescription sql.NullString
		if item.Description != "" {
			postDescription.String = item.Description
			postDescription.Valid = true
		} else {
			postDescription.Valid = false
		}

		
		createPostParams := database.CreatePostParams {
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: postDescription,
			PublishedAt: postPubDate,
			FeedID: feed.ID,
		}

		err = s.db.CreatePost(context.Background(), createPostParams)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" { /*unique violation on posts.url*/
					continue
				} 
			}
			log.Printf("could not enter %s into posts table: %w", createPostParams.Title, err)
		}

	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
	fmt.Println("===================================================================================================")
	return
}



