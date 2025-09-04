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