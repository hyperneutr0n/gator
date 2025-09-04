-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, name, url, created_at, updated_at) 
VALUES(
    DEFAULT,
    $1,
    $2,
    $3,
    DEFAULT,
    DEFAULT
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.*, u.name AS user_name
FROM feeds AS f
INNER JOIN users AS u
ON u.id = f.user_id;