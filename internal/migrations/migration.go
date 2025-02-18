package migrations

import (
	"database/sql"
	"log/slog"

	"github.com/WitnessBro/education/internal/config"
	"github.com/pressly/goose"
)

func DoMigrations(db *sql.DB) {
	config, _ := config.LoadConfig("configs/config.yaml")
	currentVersion, err := goose.GetDBVersion(db)
	if err != nil {
		slog.Error("Ошибка получения текущей версии базы данных")
	}
	migrations, err := goose.CollectMigrations(config.MigrationDir, 0, goose.MaxVersion)
	if err != nil {
		slog.Error("Ошибка получения списка миграций")
	}

	if len(migrations) > 0 && migrations[len(migrations)-1].Version > currentVersion {
		slog.Info("Обнаружены непримененные миграции. Применение...")

		if err := goose.Up(db, config.MigrationDir); err != nil {
			slog.Error("Ошибка применения миграций")
		}

		slog.Debug("Миграции успешно применены.", slog.String("path", config.MigrationDir))
	}
	slog.Info("Все миграции уже применены.")
}
