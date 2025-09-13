-- name: CreatePost :one
INSERT INTO posts (id, feed_id, title, url, description, published_at, created_at, updated_at)
VALUES(
    DEFAULT,
    $1,
    $2,
    $3,
    $4,
    $5,
    DEFAULT,
    DEFAULT
)
RETURNING *;

-- name: GetPostFromUser :many
SELECT p.*
FROM posts p
INNER JOIN feed_follows ff
ON ff.feed_id = p.feed_id
WHERE ff.user_id = $1
ORDER BY published_at DESC
LIMIT $2;