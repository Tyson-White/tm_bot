-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE invite DROP CONSTRAINT invite_group_id_fkey;
ALTER TABLE invite RENAME COLUMN group_id to groupname;
ALTER TABLE invite ALTER COLUMN groupname TYPE varchar(90);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
