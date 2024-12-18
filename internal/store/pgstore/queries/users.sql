-- name: CreateUser :one
INSERT INTO users ("username", "email", "password_hash", "bio") 
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetUserById :one
SELECT * 
FROM users
WHERE id = $1;