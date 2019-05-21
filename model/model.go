package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	DateOfBirth string `json:"DateOfBirth"`
}

func getPGConnetionStr() string {

	var pgHost string
	var pgPort string
	var pgUser string
	var pgDb string
	var pgPw string

	if len(os.Getenv("DATABASE_HOST")) == 0 {
		pgHost = "db"
	} else {
		pgHost = os.Getenv("DATABASE_HOST")
	}

	if len(os.Getenv("DATABASE_PORT")) == 0 {
		pgPort = "5432"
	} else {
		pgPort = os.Getenv("DATABASE_PORT")
	}

	if len(os.Getenv("DATABASE_USER")) == 0 {
		pgUser = "user"
	} else {
		pgUser = os.Getenv("DATABASE_USER")
	}

	if len(os.Getenv("DATABASE_DBNAME")) == 0 {
		pgDb = "user"
	} else {
		pgDb = os.Getenv("DATABASE_DBNAME")
	}

	if len(os.Getenv("DATABASE_PASSWORD")) == 0 {
		pgPw = "user"
	} else {
		pgPw = os.Getenv("DATABASE_PASSWORD")
	}

	return "host=" + pgHost + " port=" + pgPort + " user=" + pgUser + " dbname=" + pgDb + " password=" + pgPw + " sslmode=disable"

}

func DbManager() *gorm.DB {

	// db, err := gorm.Open("sqlite3", "gorm.db")
	db, err := gorm.Open("postgres", getPGConnetionStr())
	if err != nil {
		fmt.Println(err)
		panic("DB Connection Error")
	}

	return db
}

// Comment
func MigrateBD() error {

	dbConn := DbManager()

	return dbConn.AutoMigrate(User{}).Error
}

// Finding date of birth by username
func GetDateOfBirthByName(u User) string {

	dbConn := DbManager()

	dbConn.Where("name = ?", u.Name).First(&u)

	return u.DateOfBirth
}

// Check if user exists
func UserNotFoundByName(u User) bool {

	dbConn := DbManager()

	return dbConn.Where("name = ?", u.Name).First(&u).RecordNotFound()
}

// Creates new user with uniq ID
func CreateUserByName(u User) bool {

	dbConn := DbManager()

	dbConn.NewRecord(u)
	dbConn.Create(&u)

	return dbConn.NewRecord(u)
}

//Updates user's data of birth
func UpdateUserDateOfBirth(u User) string {

	dbConn := DbManager()

	dbConn.Model(&u).Where("name = ?", u.Name).Update("DateOfBirth", u.DateOfBirth)

	return u.DateOfBirth
}
