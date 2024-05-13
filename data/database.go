package data

import (
	"fmt"
	"github.com/topboyasante/achieve/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// RunMigrations uses gorm's AutoMigrate function to run migrations for all your models.
func RunMigrations(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	return err
}

// GeneratePostgresURI generated your postgres DSN using your environment keys.
func GeneratePostgresURI(c *config.Config) string {
	var (
		dbUrl    = c.DATABASE_URL
		host     = c.PG_HOST
		port     = c.PG_PORT
		dbname   = c.PG_NAME
		user     = c.PG_USER
		password = c.PG_PASS
	)
	if c.ENVIRONMENT == config.Development {
		dbUrl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	}
	return dbUrl
}

func NewPostgres(config *config.Config) (*gorm.DB, error) {
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
