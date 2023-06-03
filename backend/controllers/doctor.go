package controllers

import (
	"net/http"

	"github.com/SohamGhugare/IHP/contract"
	"github.com/gin-gonic/gin"
)

func CreateDoctor(c *gin.Context) {
	// Binding the request body
	var body struct {
		License int
		Name    string
		Phone   int
		Email   string
	}

	err := c.Bind(&body)

	// Error while binding
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create doctor, check the request body",
		})
		return
	}

	tx := contract.CreateDoctorProfile(body.License, body.Name, body.Email)
	c.JSON(http.StatusCreated, gin.H{
		"transaction_addr": tx,
	})

}
