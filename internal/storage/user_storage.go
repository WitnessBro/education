package storage

import (
	"database/sql"

	"github.com/WitnessBro/education/internal/models"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserStorage {
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
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserStorage) GetUser(id string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserStorage) CreateUser(u *models.User) error {
	err := r.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserStorage) UpdateUser(id string, u *models.User) error {
	_, err := r.db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserStorage) DeleteUser(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
