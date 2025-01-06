-- +goose Up
-- +goose StatementBegin
INSERT INTO reference (name, type)
VALUES
    ('антибиотик', 'category'),
    ('физ. pаствор', 'category'),
    ('витамин', 'category'),
    ('макро - и микро элементы', 'category'),
    ('нпвс', 'category'),
    ('гкс', 'category'),
    ('гемостатик', 'category'),
    ('противодиарейное средство', 'category'),
    ('противорвотное средство', 'category'),
    ('анальгетик', 'category'),
    ('спазмолитик', 'category');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
