-- +goose Up
-- +goose StatementBegin
ALTER TABLE preparation ADD COLUMN IF NOT EXISTS popularity INT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE preparation DROP COLUMN IF EXISTS popularity;
-- +goose StatementEnd
