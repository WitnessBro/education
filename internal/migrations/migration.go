package migrations

import (
	"database/sql"
	"log/slog"

	"github.com/pressly/goose"
)

func DoMigrations(db *sql.DB) {
	currentVersion, err := goose.GetDBVersion(db)
	if err != nil {
		slog.Error("Ошибка получения текущей версии базы данных")
	}
	migrations, err := goose.CollectMigrations("internal/migrations", 0, goose.MaxVersion)
	if err != nil {
		slog.Error("Ошибка получения списка миграций")
	}

	if len(migrations) > 0 && migrations[len(migrations)-1].Version > currentVersion {
		slog.Error("Обнаружены непримененные миграции. Применение...")

		if err := goose.Up(db, "internal/migrations"); err != nil {
			slog.Error("Ошибка применения миграций")
		}

		slog.Error("Миграции успешно применены.")
	} else {
		slog.Error("Все миграции уже применены.")
	}
	if err := goose.Up(db, "internal/migrations"); err != nil {
		slog.Error("Database already exist")
	}
}
