-- +goose Up
-- +goose StatementBegin
ALTER TABLE symptom ADD CONSTRAINT symptom_name_unique UNIQUE (name);

ALTER TABLE preparation ADD CONSTRAINT preparation_name_unique UNIQUE (name);

ALTER TABLE symptom_relation_preparation
    ADD CONSTRAINT symptom_preparation_unique UNIQUE (symptom_id, preparation_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
