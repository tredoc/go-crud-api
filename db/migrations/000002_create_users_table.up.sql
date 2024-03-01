CREATE TABLE IF NOT EXISTS users (
     id bigserial PRIMARY KEY,
     email varchar(255) NOT NULL,
     created_at timestamp DEFAULT (now()),
     password varchar(255) NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS email_index ON users ("email");