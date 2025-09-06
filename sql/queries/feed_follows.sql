-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
    VALUES(
        DEFAULT,
        $1,
        $2,
        DEFAULT,
        DEFAULT
    )
    RETURNING *
)
SELECT 
    inserted_feed_follow.*,
    users.name AS user_name,
    feeds.name AS feed_name
FROM 
    inserted_feed_follow
INNER JOIN
    users ON users.id = inserted_feed_follow.user_id
INNER JOIN
    feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowForUser :many
SELECT 
    feed_follows.*,
    users.name AS user_name,
    feeds.name AS feed_name
FROM 
    feed_follows
INNER JOIN
    users ON users.id = feed_follows.user_id
INNER JOIN
    feeds ON feeds.id = feed_follows.feed_id
WHERE
	users.id = $1;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;