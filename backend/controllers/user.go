package controllers

import (
	"net/http"

	"github.com/SohamGhugare/IHP/contract"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Binding the request body
	var body struct {
		Name string
	}

	err := c.Bind(&body)

	// Error while binding
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user, check the request body",
		})
		return
	}
	ihp, hash := contract.CreateIHPProfile("randomtest", body.Name)
	c.JSON(http.StatusCreated, gin.H{
		"ihp":              ihp,
		"transaction_hash": hash,
	})
}
