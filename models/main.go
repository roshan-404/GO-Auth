package models

import (
	"auth/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func ConnectToDb() {

	config.DB, err = gorm.Open(postgres.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		fmt.Printf("Status: %v\n", err)
	}

	config.DB.AutoMigrate(&Names{})

	fmt.Println("Database connected!")
}
