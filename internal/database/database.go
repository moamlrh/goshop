package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/moamlrh/goshop/internal/config"
)

func NewPostgresDb() (*sql.DB, error) {
	cs := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Name,
	)

	db, err := sql.Open("postgres", cs)
	if err != nil {
		return nil, err
	}

	if err := RunMigrations(db); err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
