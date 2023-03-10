-- name: CreateUser :one
INSERT INTO users (
  username, hashed_password, first_name, last_name, email
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1 OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
set username = $2,
hashed_password = $3,
first_name = $4,
last_name = $5,
email = $6
WHERE id = $1
RETURNING *;