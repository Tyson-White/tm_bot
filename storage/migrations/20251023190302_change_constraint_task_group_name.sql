-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE task ALTER COLUMN groupname TYPE varchar(90);
ALTER TABLE task ADD CONSTRAINT task_group_name_fk FOREIGN KEY (groupname) 
REFERENCES task_group(name) ON DELETE CASCADE;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE task DROP CONSTRAINT task_group_name_fk;
ALTER TABLE task ALTER COLUMN groupname TYPE int;
