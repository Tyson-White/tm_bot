-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE users_group (
	id serial PRIMARY KEY NOT NULL,
	username varchar(33) NOT NULL,
	groupname varchar(90) REFERENCES task_group(name) 	
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE users_group;
