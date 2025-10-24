-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE t_user (
	id serial PRIMARY KEY NOT NULL,
	telegram_id int UNIQUE NOT NULL,
	username varchar(33) UNIQUE NOT NULL
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE t_user;
