package models

import (
	"gorm.io/gorm"
)

type Names struct {
	gorm.Model
	Username string
	Password string 
}
