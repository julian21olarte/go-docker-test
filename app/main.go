package main

import (
	"app/main/db"
	"app/main/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	db.SetupMongo()
	println("Starting server")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World3 edited!")
	})

	e.GET("/users/create", controllers.CreateUser)
	e.GET("/users", controllers.GetUsers)
	e.Logger.Fatal(e.Start(":8080"))
}
