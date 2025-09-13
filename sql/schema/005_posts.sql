-- +goose Up
CREATE TABLE posts(
    id SERIAL PRIMARY KEY,
    feed_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    published_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', NOW()),
    created_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', NOW()),

    CONSTRAINT fk_feed
        FOREIGN KEY(feed_id)
            REFERENCES feeds(id)
            ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;