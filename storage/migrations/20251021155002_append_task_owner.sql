-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

TRUNCATE TABLE task;
ALTER TABLE task ADD COLUMN owner varchar(32) not null;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
