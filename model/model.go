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

func DbManager() *gorm.DB {

	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic("DB Connection Error")
	}

	return db
}

// Comment
func MigrateBD() error {

	dbConn := DbManager()

	return dbConn.AutoMigrate(User{}).Error
}

// Comment
func GetDateOfBirthByName(u User) string {

	dbConn := DbManager()

	dbConn.Where("name = ?", u.Name).First(&u)

	return u.DateOfBirth
}

// Comment
func UserNotFoundByName(u User) bool {

	dbConn := DbManager()

	return dbConn.Where("name = ?", u.Name).First(&u).RecordNotFound()
}

// Comment
func CreateUserByName(u User) bool {

	dbConn := DbManager()

	dbConn.NewRecord(u)
	dbConn.Create(&u)

	return dbConn.NewRecord(u)
}

// Comment
func UpdateUserDateOfBirth(u User) string {

	dbConn := DbManager()

	dbConn.Model(&u).Where("name = ?", u.Name).Update("DateOfBirth", u.DateOfBirth)

	return u.DateOfBirth
}
