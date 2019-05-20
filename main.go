package main

import (
	"./model"
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

	u.Name = c.Param("name")

	if model.UserNotFoundByName(*u) {
		return c.JSON(http.StatusOK, "User NOT found")
	} else {
		return c.JSON(http.StatusOK, model.GetDateOfBirthByName(*u))
	}
}

func putUser(c echo.Context) error {

	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	u.Name = c.Param("name")

	if model.UserNotFoundByName(*u) {
		model.CreateUserByName(*u)
	} else {
		model.UpdateUserDateOfBirth(*u)
	}

	return c.JSON(http.StatusOK, u.DateOfBirth)
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
