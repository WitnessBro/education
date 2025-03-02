package testlib

import (
	"net/http"

	"github.com/WitnessBro/education/internal/storage"
)

type TestApp struct {
	StorageTask       *storage.TaskStorage
	StorageUser       *storage.UserStorage
	TaskManagerClient *http.Client
}
