package db

import (
    "errors"
    "financial-manager-api/utils/logger"
    "fmt"
    "os"

    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {
    databaseUrl := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=disable",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    m, err := migrate.New("file://migrations", databaseUrl)
    if err != nil {
        return fmt.Errorf("erro ao criar migrator: %w", err)
    }

    defer func(m *migrate.Migrate) {
        sourceErr, dbErr := m.Close()
        if sourceErr != nil {
            logger.L.Warnw("erro ao fechar source das migrations", "error", sourceErr)
        }
        if dbErr != nil {
            logger.L.Warnw("erro ao fechar conexão das migrations", "error", dbErr)
        }
    }(m)

    if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
        return fmt.Errorf("erro ao rodar migrations: %w", err)
    }

    return nil
}
