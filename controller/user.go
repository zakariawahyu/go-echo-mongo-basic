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

	result, errGet := model.GetALlUser(ctx)
	if errGet != nil {
		return errGet
	}

	defer result.Close(ctx)
	for result.Next(ctx) {
		var user entity.User
		if errDecode := result.Decode(&user); errDecode != nil {
			return errDecode
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, response.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Results: users,
	})
}

func CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user entity.User
	defer cancel()

	// validate the request body
	if errBind := c.Bind(&user); errBind != nil {
		return errBind
	}

	//use the validator to validate request
	if validateErr := validate.Struct(&user); validateErr != nil {
		return validateErr
	}

	newUser := entity.User{
		Id:        primitive.NewObjectID(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Location:  user.Location,
		Title:     user.Title,
	}

	result, errCreate := model.CreateUser(ctx, newUser)
	if errCreate != nil {
		return errCreate
	}

	return c.JSON(http.StatusCreated, response.WebResponse{
		Code:   http.StatusCreated,
		Status: "success",
		Results: &echo.Map{
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
		return errObj
	}

	if errGet := model.GetUserById(ctx, objectId, &user); errGet != nil {
		return errGet
	}

	return c.JSON(http.StatusOK, response.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Results: user,
	})
}

func UpdateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var user entity.User
	defer cancel()

	objectId, errObj := primitive.ObjectIDFromHex(id)
	if errObj != nil {
		return errObj
	}

	//validate request body
	if errBind := c.Bind(&user); errBind != nil {
		return errBind
	}

	if errGet := model.GetUserById(ctx, objectId, &entity.User{}); errGet != nil {
		return errGet
	}

	//validation
	if validateErr := validate.Struct(&user); validateErr != nil {
		return validateErr
	}

	updateUser := bson.M{"firstname": user.FirstName, "lastname": user.LastName, "username": user.Username, "location": user.Location, "title": user.Title}

	result, errUpdate := model.UpdateUser(ctx, objectId, updateUser)
	if errUpdate != nil {
		return errUpdate
	}

	return c.JSON(http.StatusOK, response.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Results: result,
	})
}

func DeleteUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	objectID, errObj := primitive.ObjectIDFromHex(id)
	if errObj != nil {
		return errObj
	}

	if errGet := model.GetUserById(ctx, objectID, &entity.User{}); errGet != nil {
		return errGet
	}

	result, errDelete := model.DeleteUser(ctx, objectID)
	if errDelete != nil {
		return errDelete
	}

	return c.JSON(http.StatusOK, response.WebResponse{
		Code:    http.StatusOK,
		Status:  "error",
		Results: result,
	})
}
