package helper

import (
	"regexp"
)

// Comment
//func DateOfBirthValid(dateStrung String) bool {
//
//	dbConn := DbManager()
//
//	return dbConn.AutoMigrate(User{}).Error
//}

// Comment
func NameValid(name string) bool {

	matched, err := regexp.MatchString(`(^+[a-zA-Z]+$)`, name)
	if err != nil {
		matched = false
	}

	return matched
}
