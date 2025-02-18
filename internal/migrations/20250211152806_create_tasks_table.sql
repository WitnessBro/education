-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL PRIMARY KEY, 
	task_description TEXT, 
	task_status TEXT,
	title TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	user_id BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
