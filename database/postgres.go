package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct{
	host		string
	port		string
	password	string
	user		string
	dbName		string
	sslMode		string
}

func establishConnection(config *Config)(*gorm.DB, error){
	dsn := fmt.Sprintf("host=%s port=%s password=%s dbname=%s sslmode=%s", config.host, config.port, config.password, config.dbName, config.sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}