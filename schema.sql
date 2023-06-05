DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id BIGSERIAL PRIMARY KEY,
    author_id BIGINT REFERENCES authors(id) DEFAULT 0,
    title TEXT NOT NULL,
    content TEXT,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now())
);

INSERT INTO authors(id, name) VALUES(0, 'default');