-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at)
VALUES (
    DEFAULT,
    $1, 
    DEFAULT, 
    DEFAULT
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE name=$1;

-- name: ResetUser :exec
DELETE FROM users;