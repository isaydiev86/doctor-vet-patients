-- +goose Up
-- +goose StatementBegin
INSERT INTO symptom_relation_preparation (symptom_id, preparation_id)
VALUES
    (1, 1),
    (2, 1),
    (3, 1),
    (4, 1),
    (5, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
