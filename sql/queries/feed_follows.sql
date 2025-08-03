-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT inserted_feed_follow.*, users.name AS userName, feeds.name AS feedName
FROM inserted_feed_follow
INNER JOIN users ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds on feeds.id = inserted_feed_follow.feed_id;


-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feedName
FROM feed_follows
INNER JOIN feeds on feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;


-- name: DeleteFeedFollowForUser :exec
DELETE FROM feed_follows 
WHERE feed_follows.feed_id = (
    SELECT feeds.id FROM feeds WHERE feeds.url = $1
)
AND
feed_follows.user_id = $2;