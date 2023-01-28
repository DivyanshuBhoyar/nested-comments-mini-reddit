BEGIN;

-- a new postgres schema for posts table with author id
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY ,
    author_email VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY  (author_email) references users(email) ON DELETE CASCADE
);

COMMIT;