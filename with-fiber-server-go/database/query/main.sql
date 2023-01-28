-- name: CreatePost :one
INSERT INTO posts (title, body, author_email)
VALUES ($1, $2, $3) RETURNING * ;


-- name: CreateComment :one
INSERT INTO comments (content, post_id, user_email, parent_id)
VALUES ($1, $2, $3, $4) RETURNING * ;

-- name: GetPostFeed :many
SELECT p.id, p.title, p.author_email, p.created_at, u.email
FROM posts p
INNER JOIN users u ON p.author_email = u.email
ORDER BY p.created_at DESC;


-- name: GetPost :one
SELECT p.title, p.body, p.created_at, users.name as userame
FROM  posts as p
JOIN  users ON p.author_email = users.email
WHERE p.id = $1;

-- // get all comments for a post
-- name: GetComments :many
SELECT c.id, c.content, c.created_at, c.parent_id, u.name as username
FROM comments c
JOIN users u ON c.user_email = u.email
WHERE c.post_id = $1;
