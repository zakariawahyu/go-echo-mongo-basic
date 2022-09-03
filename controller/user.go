package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-mongo-basic/config"
	"github.com/zakariawahyu/go-echo-mongo-basic/model"
	"github.com/zakariawahyu/go-echo-mongo-basic/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var validate = validator.New()

func CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.User
	defer cancel()

	// validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Result:  err.Error(),
		})
	}

	// use the validator to validate request
	if validateErr := validate.Struct(&user); validateErr != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Result:  validateErr.Error(),
		})
	}

	newUser := model.User{
		Id:        primitive.NewObjectID(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Location:  user.Location,
		Title:     user.Title,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Result:  err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.UserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Result: &echo.Map{
			"data": result,
		}})
}

func GetUserById(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var user model.User
	defer cancel()

	objectId, errObj := primitive.ObjectIDFromHex(id)
	if errObj != nil {
		return c.JSON(http.StatusInternalServerError, response.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Result:  errObj.Error(),
		})
	}

	if err := userCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&user); err != nil {
		return c.JSON(http.StatusNotFound, response.UserResponse{
			Status:  http.StatusNotFound,
			Message: "error",
			Result:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  user,
	})
}
