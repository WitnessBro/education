package testlib

import (
	"net/http"

	"github.com/WitnessBro/education/internal/config"
	"github.com/WitnessBro/education/internal/storage"
	"github.com/WitnessBro/education/pkg/db"
)

func TestInitApp(config *config.Config) *TestApp {
	db, _ := db.Connect(config.DatabaseURL)
	userRepo := storage.NewUserStorage(db)
	taskRepo := storage.NewTaskStorage(db)

	return &TestApp{
		StorageTask:       taskRepo,
		StorageUser:       userRepo,
		TaskManagerClient: &http.Client{},
	}
}
