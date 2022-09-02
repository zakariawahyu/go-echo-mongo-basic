package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Hello world form Echo and MongoDB",
		})
	})

	e.Logger.Fatal(e.Start("localhost:8081"))
}
