-- name: CreateUser :one
INSERT INTO users (
  username, hashed_password, first_name, last_name, email
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUsers :many
SELECT * FROM users;