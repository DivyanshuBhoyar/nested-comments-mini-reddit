-- name: CreateUser :one
INSERT INTO users (email, name, password_hash)
VALUES ($1, $2, $3) RETURNING email, name, created_at ;

-- name: GetUser :one
SELECT password_hash  FROM users WHERE email = $1 ;