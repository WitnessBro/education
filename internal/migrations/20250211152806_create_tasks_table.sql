-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL PRIMARY KEY, 
	description TEXT, 
	status TEXT,
	title TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
