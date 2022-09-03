package main

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-mongo-basic/config"
	"github.com/zakariawahyu/go-echo-mongo-basic/routes"
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

	// Define Routes
	routes.UserRoutes(e)

	e.Logger.Fatal(e.Start("localhost:8081"))
}
