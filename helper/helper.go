package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func dateParser(dateString string) (time.Time, time.Time) {

	const layout = "2006-01-02"

	currentDay, err := time.Parse(layout, time.Now().Format(layout))
	if err != nil {
		fmt.Println(err)
	}

	parsedDay, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println(err)
	}

	return currentDay, parsedDay
}

// Comment
func BirthDayToday(dateString string) bool {

	currentDay, wasBorn := dateParser(dateString)

	return wasBorn.Month() == currentDay.Month() && wasBorn.Day() == currentDay.Day()
}

// Comment
func BirthDateInThePast(dateString string) bool {

	currentDay, wasBorn := dateParser(dateString)

	return currentDay.After(wasBorn)
}

// Comment
func GetDaysBeforeBirthday(dateString string) string {

	var diff float64
	var wasBornMonthString string

	currentDay, wasBorn := dateParser(dateString)

	if int(wasBorn.Month()) <= 9 {
		wasBornMonthString = "0" + strconv.Itoa(int(wasBorn.Month()))
	} else {
		wasBornMonthString = strconv.Itoa(int(wasBorn.Month()))
	}

	thisYearBirthdayDateString := strconv.Itoa(currentDay.Year()) + "-" + wasBornMonthString + "-" + strconv.Itoa(wasBorn.Day())
	nextYearBirthdayDateString := strconv.Itoa(currentDay.Year()+1) + "-" + wasBornMonthString + "-" + strconv.Itoa(wasBorn.Day())

	_, thisYearBirthdayDate := dateParser(thisYearBirthdayDateString)
	_, nextYearBirthdayDate := dateParser(nextYearBirthdayDateString)

	if currentDay.Before(thisYearBirthdayDate) {
		diff = thisYearBirthdayDate.Sub(currentDay).Hours()
	} else {
		diff = nextYearBirthdayDate.Sub(currentDay).Hours()
	}

	return strconv.FormatFloat(diff/24, 'f', 0, 64)
}

// Comment
func NameValid(name string) bool {

	matched, err := regexp.MatchString(`(^+[a-zA-Z]+$)`, name)
	if err != nil {
		matched = false
	}

	return matched
}

func BirthDateValid(dateString string) bool {

	const layout = "2006-01-02"

	var parsed bool

	_, err := time.Parse(layout, dateString)
	if err != nil {
		parsed = false
	} else {
		parsed = true
	}

	matched, err := regexp.MatchString(`(^[0-9]{4}-[0-1]{1}[0-9]{1}-[0-3]{1}[0-9]{1}$)`, dateString)
	if err != nil {
		matched = false
	}

	return matched && parsed
}
