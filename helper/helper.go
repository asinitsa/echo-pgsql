package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func dateFormatter(dateString string) (time.Time, time.Time) {

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

	currentDay, wasBorn := dateFormatter(dateString)

	return wasBorn.Month() == currentDay.Month() && wasBorn.Day() == currentDay.Day()
}

// Comment
func BirthDateInThePast(dateString string) bool {

	currentDay, wasBorn := dateFormatter(dateString)

	return currentDay.After(wasBorn)
}

// Comment
func GetDaysBeforeBirthday(dateString string) string {

	var diff float64
	var wasBornMonthString string

	currentDay, wasBorn := dateFormatter(dateString)

	if int(wasBorn.Month()) <= 9 {
		wasBornMonthString = "0" + strconv.Itoa(int(wasBorn.Month()))
	} else {
		wasBornMonthString = strconv.Itoa(int(wasBorn.Month()))
	}

	thisYearBirthdayDateString := strconv.Itoa(currentDay.Year()) + "-" + wasBornMonthString + "-" + strconv.Itoa(wasBorn.Day())
	nextYearBirthdayDateString := strconv.Itoa(currentDay.Year()+1) + "-" + wasBornMonthString + "-" + strconv.Itoa(wasBorn.Day())

	_, thisYearBirthdayDate := dateFormatter(thisYearBirthdayDateString)
	_, nextYearBirthdayDate := dateFormatter(nextYearBirthdayDateString)

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
