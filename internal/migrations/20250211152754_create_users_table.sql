-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY, 
	name TEXT, 
	email TEXT, 
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
	status TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
