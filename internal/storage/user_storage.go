package storage

import (
	"database/sql"

	"github.com/WitnessBro/education/internal/models"
)

type UserStorage struct {
	db *sql.DB
}

// TODO вместе с InitConnection поменять тут

// TODO подумать про ауторизэйшн токен для получения всех тасок
func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (r *UserStorage) GetUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.Status); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserStorage) GetUser(id int) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow("SELECT id, name, email, created_at, status FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.Status)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserStorage) CreateUser(u *models.User) error {
	err := r.db.QueryRow("INSERT INTO users (name, email, status) VALUES ($1, $2, $3) RETURNING id", u.Name, u.Email, u.Status).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserStorage) UpdateUser(id int, u *models.User) error {
	_, err := r.db.Exec("UPDATE users SET name = $1, email = $2, status = $3 WHERE id = $4", u.Name, u.Email, u.Status, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserStorage) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
