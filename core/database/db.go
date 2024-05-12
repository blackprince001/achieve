package database

import (
	"fmt"
	"github.com/topboyasante/achieve/core"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// RunMigrations uses gorm's AutoMigrate function to run migrations for all your models.
func RunMigrations(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	return err
}

// GeneratePostgresURI generated your postgres DSN using your environment keys.
func GeneratePostgresURI(config *core.Config) string {
	var (
		dbUrl    = config.DATABASE_URL
		host     = config.PG_HOST
		port     = config.PG_PORT
		dbname   = config.PG_NAME
		user     = config.PG_USER
		password = config.PG_PASS
	)
	if config.ENVIRONMENT == core.Development {
		dbUrl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	}
	return dbUrl
}


func NewPostgres(config *core.Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	db, err = gorm.Open(postgres.Open(GeneratePostgresURI(config)), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db = db.Debug()

	return db, nil
}
