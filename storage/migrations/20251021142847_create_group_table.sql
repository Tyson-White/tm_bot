-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE task_group (
	id serial PRIMARY KEY not null,
	name varchar(90) not null,
	creator varchar(32) not null
);

ALTER TABLE task ADD COLUMN group_id int REFERENCES task_group(id) ON DELETE CASCADE;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE task DROP COLUMN group_id;
DROP TABLE task_group;


