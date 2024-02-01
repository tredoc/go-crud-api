CREATE TABLE IF NOT EXISTS books (
  id bigserial PRIMARY KEY,
  title varchar(255) NOT NULL,
  publish_date date NOT NULL,
  created_at timestamp DEFAULT (now()),
  ISBN varchar(100) NOT NULL,
  pages smallint NOT NULL
);

CREATE TABLE IF NOT EXISTS authors (
  id bigserial PRIMARY KEY,
  first_name varchar(100) NOT NULL,
  middle_name varchar(100),
  last_name varchar(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS genres (
  id bigserial PRIMARY KEY,
  name varchar(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS book_genre (
  book_id bigint NOT NULL,
  genre_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS book_author (
  book_id bigint NOT NULL,
  author_id bigint NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS book_genre_index ON book_genre ("book_id", "genre_id");
CREATE UNIQUE INDEX IF NOT EXISTS book_author_index ON book_author ("book_id", "author_id");

ALTER TABLE book_genre ADD FOREIGN KEY ("book_id") REFERENCES books(id);
ALTER TABLE book_genre ADD FOREIGN KEY ("genre_id") REFERENCES genres(id);
ALTER TABLE book_author ADD FOREIGN KEY ("book_id") REFERENCES books(id);
ALTER TABLE book_author ADD FOREIGN KEY ("author_id") REFERENCES authors(id);
