-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE task (
	id serial PRIMARY KEY not null,
	title varchar(90) not null,
	description varchar(300),
	created_at timestamp default now()
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
