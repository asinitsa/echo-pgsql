package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_b"`
}
