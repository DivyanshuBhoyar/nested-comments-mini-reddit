// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package database

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, name, password_hash)
VALUES ($1, $2, $3) RETURNING email, name, created_at
`

type CreateUserParams struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
}

type CreateUserRow struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.Email, arg.Name, arg.PasswordHash)
	var i CreateUserRow
	err := row.Scan(&i.Email, &i.Name, &i.CreatedAt)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT password_hash  FROM users WHERE email = $1
`

func (q *Queries) GetUser(ctx context.Context, email string) (string, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, email)
	var password_hash string
	err := row.Scan(&password_hash)
	return password_hash, err
}
