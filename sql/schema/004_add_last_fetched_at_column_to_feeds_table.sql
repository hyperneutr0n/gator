-- +goose Up
ALTER TABLE feeds
ADD COLUMN last_fetched_at TIMESTAMP NULL DEFAULT TIMEZONE('utc', NOW());

-- +goose Down
ALTER TABLE feeds
DROP COLUMN last_fetched_at;