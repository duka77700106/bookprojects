package config

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func SetupDatabaseConnection() (*gorm.DB, error) {
	host := "db"
	user := "bookuser"
	password := "bookpassword"
	dbname := "bookdb"
	port := "5432"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("SQL connection error:", err)
		return nil, err
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Println("Migration driver error:", err)
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/config/migration",
		"postgres", driver)
	if err != nil {
		log.Println("Migration error:", err)
		return nil, err
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Println("Migration up error:", err)
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return gormDB, nil
}
