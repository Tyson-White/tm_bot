-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE invite RENAME COLUMN creator_id TO creator;
ALTER TABLE invite RENAME COLUMN invited_id TO invited;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd


ALTER TABLE invite RENAME COLUMN creator TO creator_id ;
ALTER TABLE invite RENAME COLUMN invited TO invited_id;
