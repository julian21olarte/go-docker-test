package main

import (
	"app/main/controllers/users"
	"app/main/db"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	db.SetupMongo()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World3 edited!")
	})

	v1 := e.Group("/v1")
	users.Routes(v1)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
