-- +goose Up
CREATE TABLE feeds (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', NOW()),

    CONSTRAINT fk_user
        FOREIGN KEY(user_id) 
	        REFERENCES users(id)
	        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;