package controllers

import (
	"context"
	"echo-mongo/configs"
	"echo-mongo/models"
	"echo-mongo/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status: http.StatusBadRequest, 
			Message: "error", 
			Data: &echo.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status: http.StatusBadRequest,
			Message: "error",
			Data: &echo.Map{"data": validationErr.Error()},
		})
	}

	newUser := models.User{
		Id: primitive.NewObjectID(),
		Name: user.Name,
		Location: user.Location,
		Title: user.Title,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{
			Status: http.StatusBadRequest,
			Message: "error",
			Data: &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusCreated, responses.UserResponse{
		Status: http.StatusCreated,
		Message: "success",
		Data: &echo.Map{"data": result},
	})
}