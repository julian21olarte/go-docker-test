package controllers

import (
	"app/main/db"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

const usersCollection = "users"

// GetUsers func
func GetUsers(c echo.Context) error {
	users := []map[string]interface{}{}
	if db.MongoClient != nil {
		cursor, err := db.MongoClient.
			Database(db.MongoDatabase).
			Collection(usersCollection).
			Find(context.TODO(), bson.D{})
		if err != nil {
			c.Error(err)
		}

		err = cursor.All(context.TODO(), &users)
		if err != nil {
			println(err.Error())
			c.Error(err)
		}
		return c.JSON(http.StatusOK, users)
	}
	return c.JSON(http.StatusInternalServerError, "Error on mongo connection")
}

// CreateUser func
func CreateUser(c echo.Context) error {
	user := map[string]interface{}{
		"name":  c.QueryParam("name"),
		"age":   c.QueryParam("age"),
		"phone": c.QueryParam("phone"),
	}
	if db.MongoClient != nil {
		_, err := db.MongoClient.
			Database(db.MongoDatabase).
			Collection(usersCollection).
			InsertOne(context.TODO(), user)
		if err != nil {
			c.Error(err)
		}

		return c.JSON(http.StatusOK, "ok")
	}
	return c.JSON(http.StatusInternalServerError, "Error on mongo connection")
}