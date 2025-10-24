-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE task DROP CONSTRAINT task_group_id_fkey;
ALTER TABLE task RENAME COLUMN group_id TO groupname;


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE task RENAME COLUMN groupname TO group_id;
ALTER TABLE task ADD CONSTRAINT task_group_id_fkey UNIQUE (task_group);

