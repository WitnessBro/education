package storage

import (
	"database/sql"

	"github.com/WitnessBro/education/internal/models"
)

type TaskStorage struct {
	db *sql.DB
}

// TODO вместе с InitConnection поменять тут
func NewTaskStorage(db *sql.DB) *TaskStorage {
	return &TaskStorage{db: db}
}

// Index on UserId

// TODO доделать
func (r *TaskStorage) GetUsersGetTasksId(id int) ([]models.Task, error) {
	rows, err := r.db.Query("SELECT title, task_description, task_status, tasks.created_at, user_id from tasks join users on tasks.user_id = users.id WHERE users.id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.Title, &t.TaskDescription, &t.TaskStatus, &t.CreatedAt, &t.UserId); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskStorage) GetTaskById(id string) (*models.Task, error) {
	var t models.Task
	err := r.db.QueryRow("SELECT title, task_description, task_status, created_at, user_id FROM tasks WHERE id = $1", id).Scan(&t.Title, &t.TaskDescription, &t.TaskStatus, &t.CreatedAt, &t.UserId)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TaskStorage) CreateTask(t *models.Task) error {
	err := r.db.QueryRow("INSERT INTO tasks (title, task_description, task_status, user_id) VALUES ($1, $2, $3, $4) RETURNING id", t.Title, t.TaskDescription, t.TaskStatus, t.UserId).Scan(&t.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskStorage) UpdateTask(id string, t *models.Task) error {
	_, err := r.db.Exec("UPDATE tasks SET title = $1, description = $2, task_status = $3 WHERE id = $4", t.Title, t.TaskDescription, t.TaskStatus, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskStorage) DeleteTask(id string) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
