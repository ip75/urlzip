package repository

import (
	"github.com/ip75/urlzip/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// UrlZipDatabase represents a database
type UrlZipDatabase struct {
	*gorm.DB
}

// Connect to the database
func Connect(cfg config.Config) (IUZRepository, error) {

	db, err := gorm.Open(postgres.Open(cfg.ComposeDSN()), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}
	postgresDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	postgresDb.SetMaxIdleConns(10)
	postgresDb.SetMaxOpenConns(100)

	return &UrlZipDatabase{db}, nil
}
