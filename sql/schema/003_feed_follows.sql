-- +goose Up
CREATE TABLE feed_follows(
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    feed_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', NOW()),

    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    
    CONSTRAINT fk_feed
        FOREIGN KEY(feed_id)
        REFERENCES feeds(id)
        ON DELETE CASCADE,

    UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;