-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id          BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    user_id     uuid NOT NULL,
    fio         VARCHAR(255),
    role        VARCHAR(20)
);
COMMENT ON TABLE users IS 'Пользователь';
COMMENT ON COLUMN users.id IS 'ID записи';
COMMENT ON COLUMN users.user_id IS 'ID в keycloak';
COMMENT ON COLUMN users.role IS 'роль  в keycloak';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
