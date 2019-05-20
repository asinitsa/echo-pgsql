package main

import (
	"./model"
	"github.com/jinzhu/gorm"
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

	DateOfBirth := model.GetDateOfBirthByName(*u)

	return c.JSON(http.StatusOK, DateOfBirth)
}

func putUser(c echo.Context) error {

	dbConn, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic("DB Connection Error")
	}

	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	u.Name = c.Param("name")

	dbConn.NewRecord(u)
	dbConn.Create(&u)
	dbConn.NewRecord(u)

	defer dbConn.Close()
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
