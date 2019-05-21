package main

import (
	"github.com/asinitsa/echo-pgsql/helper"
	"github.com/asinitsa/echo-pgsql/model"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func getUser(c echo.Context) error {

	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if helper.NameValid(c.Param("name")) {
		u.Name = c.Param("name")
	} else {
		return c.JSON(http.StatusBadRequest, "Username must contain only letters")
	}

	if model.UserNotFoundByName(*u) {
		return c.JSON(http.StatusNotFound, "Username NOT found")
	}

	if helper.BirthDayToday(model.GetDateOfBirthByName(*u)) {
		return c.JSON(http.StatusOK, "Hello "+u.Name+"! Happy birthday!")
	}

	N := helper.GetDaysBeforeBirthday(model.GetDateOfBirthByName(*u))

	return c.JSON(http.StatusOK, "Your birthday is in "+N+" days(s) "+model.GetDateOfBirthByName(*u))
}

func putUser(c echo.Context) error {

	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if helper.NameValid(c.Param("name")) {
		u.Name = c.Param("name")
	} else {
		return c.JSON(http.StatusBadRequest, "Username must contain only letters!")
	}

	if !helper.BirthDateValid(u.DateOfBirth) {
		return c.JSON(http.StatusBadRequest, "DateOfBirth must be formatted: { 'DateOfBirth': 'YYYY-MM-DD' }")
	}

	if !helper.BirthDateInThePast(u.DateOfBirth) {
		return c.JSON(http.StatusBadRequest, "Birthday date must be in the past!")
	}

	if model.UserNotFoundByName(*u) {
		model.CreateUserByName(*u)
	} else {
		model.UpdateUserDateOfBirth(*u)
	}

	return c.JSON(http.StatusNoContent, "OK")
}

func main() {

	err := model.MigrateBD()
	if err != nil {
		panic("DB Migration Error")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello/:name", getUser)
	e.PUT("/hello/:name", putUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
