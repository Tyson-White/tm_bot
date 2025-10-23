-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE invite (
	id serial PRIMARY KEY not null,
	group_id int REFERENCES task_group(id) ON DELETE CASCADE,
	creator_id varchar(32),
	invited_id varchar(32)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE invite;
