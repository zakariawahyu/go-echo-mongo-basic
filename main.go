package main

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-mongo-basic/config"
	"net/http"
)

func main() {
	e := echo.New()

	//Run Mongo
	config.ConnectDB()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Hello world form Echo and MongoDB",
		})
	})

	e.Logger.Fatal(e.Start("localhost:8081"))
}
