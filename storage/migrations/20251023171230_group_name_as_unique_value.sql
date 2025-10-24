-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE task_group
ADD CONSTRAINT task_group_name_key UNIQUE (name);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE task_group DROP CONSTRAINT task_group_name_key;
