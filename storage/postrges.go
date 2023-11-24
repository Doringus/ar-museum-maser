package storage

import (
	"ar-museum-backend/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreConfig struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSL string
}

var (
	DBCon *gorm.DB
)

func CreateConnection(config *PostgreConfig)(*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSL,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	err = db.AutoMigrate(models.Client{}, models.Exhibit{}, models.Exhibition{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
