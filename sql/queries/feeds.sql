-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, name, url, created_at, updated_at, last_fetched_at) 
VALUES(
    DEFAULT,
    $1,
    $2,
    $3,
    DEFAULT,
    DEFAULT,
    NULL
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.*, u.name AS user_name
FROM feeds AS f
INNER JOIN users AS u
ON u.id = f.user_id;

-- name: GetFeed :one
SELECT f.*, u.name AS user_name
FROM feeds AS f
INNER JOIN users AS u
ON u.id = f.user_id
WHERE f.url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds 
SET updated_at = DEFAULT, last_fetched_at = DEFAULT
WHERE id = $1;

-- name: GetNextFeedToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST;