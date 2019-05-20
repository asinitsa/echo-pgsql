package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	DateOfBirth string `json:"DateOfBirth"`
}

// Comment
func MigrateBD() error {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic("DB Connection Error")
	}

	er := db.AutoMigrate(User{}).Error

	defer db.Close()

	return er
}

// Comment
func GetDateOfBirthByName(u User) string {
	dbConn, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic("DB Connection Error")
	}

	dbConn.Where("name = ?", u.Name).First(&u)

	defer dbConn.Close()

	return u.DateOfBirth
}
