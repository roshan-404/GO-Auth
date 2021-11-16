package config

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     5342,
		User:     "postgres",
		Password: "",
		DBName:   "auth",
	}
	return &dbConfig
}

func DbURL(db *DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable ", db.Host, db.User, db.Password, db.DBName, db.Port)
}
