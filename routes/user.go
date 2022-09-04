package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-mongo-basic/controller"
)

func UserRoutes(e *echo.Echo) {
	e.GET("/user", controller.GetAllUser)
	e.POST("/user", controller.CreateUser)
	e.GET("/user/:id", controller.GetUserById)
	e.PUT("/user/:id", controller.UpdateUser)
	e.DELETE("/user/:id", controller.DeleteUser)
}
