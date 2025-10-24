-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

TRUNCATE TABLE invite;
ALTER TABLE invite ADD CONSTRAINT invite_group_name_fk FOREIGN KEY (groupname) 
REFERENCES task_group(name) ON DELETE CASCADE;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
