package db

import (
	"fmt"
	"github.com/dbl90/airfreight/internal/models"
	"github.com/dbl90/airfreight/internal/models/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	database *gorm.DB
}

func NewDBClient(config *config.DbConfig) (*DBClient, error) {
	// Connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Berlin",
		config.Host, config.User, config.Password, config.DBName, config.Port)
	database := postgres.Open(dsn)
	db, err := gorm.Open(database, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DBClient{database: db}, nil
}

func (db *DBClient) GetDatabase() *gorm.DB {
	return db.database
}

func (db *DBClient) Close() error {
	sqlDB, err := db.database.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *DBClient) Insert(thing interface{}) *gorm.DB {
	return db.database.Create(thing)
}

func (db *DBClient) GetMawbByNumber(number string) *models.Mawb {
	var mawb models.Mawb
	db.database.Where(&models.Mawb{Number: number}).Find(&mawb)
	return &mawb
}

func (db *DBClient) GetHawbByNumber(number string) models.Hawb {
	var hawb models.Hawb
	db.database.Where(&models.Hawb{Number: number}).Find(&hawb)
	return hawb
}
