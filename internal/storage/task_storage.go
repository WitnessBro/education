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
func (r *TaskStorage) GetTasks() ([]models.Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status); err != nil {
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
	err := r.db.QueryRow("SELECT * FROM tasks WHERE id = $1", id).Scan(&t.Id, &t.Title, &t.Description, &t.Status)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TaskStorage) CreateTask(t *models.Task) error {
	err := r.db.QueryRow("INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id", t.Title, t.Description, t.Status).Scan(&t.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskStorage) UpdateTask(id int, t *models.Task) error {
	_, err := r.db.Exec("UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4", t.Title, t.Description, t.Status, id)
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
