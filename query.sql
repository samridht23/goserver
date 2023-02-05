-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO users (name,bio)VALUES($1,$2) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
