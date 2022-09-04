package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-mongo-basic/config"
	"github.com/zakariawahyu/go-echo-mongo-basic/entity"
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

func GetAllUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []entity.User
	defer cancel()

	result, err := model.GetALlUser(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Result:  err.Error(),
		})
	}

	defer result.Close(ctx)
	for result.Next(ctx) {
		var user entity.User
		if err = result.Decode(&user); err != nil {
			return c.JSON(http.StatusInternalServerError, response.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Result:  err.Error(),
			})
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, response.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  users,
	})
}

func CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user entity.User
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

	newUser := entity.User{
		Id:        primitive.NewObjectID(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Location:  user.Location,
		Title:     user.Title,
	}

	result, err := model.CreateUser(ctx, newUser)
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
	var user entity.User
	defer cancel()

	objectId, errObj := primitive.ObjectIDFromHex(id)
	if errObj != nil {
		return c.JSON(http.StatusInternalServerError, response.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Result:  errObj.Error(),
		})
	}

	if err := model.GetUserById(ctx, objectId, &user); err != nil {
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

func UpdateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var user entity.User
	defer cancel()

	objectId, errObj := primitive.ObjectIDFromHex(id)
	if errObj != nil {
		return c.JSON(http.StatusInternalServerError, response.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Result:  errObj.Error(),
		})
	}

	//validate request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Result:  err.Error(),
		})
	}

	if err := model.GetUserById(ctx, objectId, &entity.User{}); err != nil {
		return c.JSON(http.StatusNotFound, response.UserResponse{
			Status:  http.StatusNotFound,
			Message: "error",
			Result:  err.Error(),
		})
	}

	//validation
	if validateErr := validate.Struct(&user); validateErr != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Result:  validateErr.Error(),
		})
	}

	updateUser := bson.M{"firstname": user.FirstName, "lastname": user.LastName, "username": user.Username, "location": user.Location, "title": user.Title}

	result, err := model.UpdateUser(ctx, objectId, updateUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Result:  errObj.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  result,
	})
}

func DeleteUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	objectID, errObj := primitive.ObjectIDFromHex(id)
	if errObj != nil {
		return c.JSON(http.StatusInternalServerError, response.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Result:  errObj.Error(),
		})
	}

	if err := model.GetUserById(ctx, objectID, &entity.User{}); err != nil {
		return c.JSON(http.StatusNotFound, response.UserResponse{
			Status:  http.StatusNotFound,
			Message: "error",
			Result:  err.Error(),
		})
	}

	result, err := model.DeleteUser(ctx, objectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Result:  errObj.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.UserResponse{
		Status:  http.StatusOK,
		Message: "error",
		Result:  result,
	})
}
