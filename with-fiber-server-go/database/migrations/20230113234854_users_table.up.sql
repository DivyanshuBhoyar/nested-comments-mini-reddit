BEGIN;

CREATE TABLE IF NOT EXISTS users(
    name VARCHAR(30) NOT NULL ,
    email VARCHAR(255) NOT NULL PRIMARY KEY,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

COMMIT;