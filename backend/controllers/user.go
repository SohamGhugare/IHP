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

func GetUser(c *gin.Context) {
	var body struct {
		IHP int
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ihp",
		})
		return
	}
	uri, name := contract.GetIHPProfile(body.IHP)
	c.JSON(http.StatusFound, gin.H{
		"name": name,
		"uri":  uri,
	})
}
