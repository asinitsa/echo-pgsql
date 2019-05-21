package helper

import (
	"fmt"
	"regexp"
	"time"
)

// Comment
func BirthDayToday(dateStrung string) bool {

	const layout = "2006-01-02"

	now := time.Now()

	t, err := time.Parse(layout, dateStrung)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t)
	fmt.Println(now)

	return t.Month() == now.Month() && t.Day() == now.Day()
}

// Comment
func BirthDayNotInThePast(dateStrung string) bool {

	const layout = "2006-01-02"

	now := time.Now()

	t, err := time.Parse(layout, dateStrung)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t)
	fmt.Println(now)

	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

// Comment
func NameValid(name string) bool {

	matched, err := regexp.MatchString(`(^+[a-zA-Z]+$)`, name)
	if err != nil {
		matched = false
	}

	return matched
}
