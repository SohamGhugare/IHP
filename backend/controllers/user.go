package controllers

import (
	"net/http"

	"github.com/SohamGhugare/IHP/database"
	"github.com/SohamGhugare/IHP/models"
	"github.com/SohamGhugare/IHP/utility"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	// Binding the request body
	var body struct {
		Name     string
		Email    string
		Password string
	}

	err := c.Bind(&body)

	// Error while binding
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user, check the request body",
		})
		return
	}

	// Hashing the password
	hash, err := utility.GenerateHash(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password, check server logs",
		})
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: hash,
	}

	database.CreateUser(user)
}
