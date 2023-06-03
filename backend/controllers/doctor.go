package controllers

import (
	"net/http"
	"time"

	"github.com/SohamGhugare/IHP/contract"
	"github.com/SohamGhugare/IHP/utility"
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

	pin, tx := contract.CreateDoctorProfile(body.License, body.Name, body.Email)
	c.JSON(http.StatusCreated, gin.H{
		"transaction_hash": tx,
		"pin":              pin,
	})

}

func LoginDoctor(c *gin.Context) {
	// Binding the request body
	var body struct {
		License int
		Pin     int
	}

	err := c.Bind(&body)

	// Error while binding
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create doctor, check the request body",
		})
		return
	}

	name, email, pin := contract.GetDoctorProfile(body.License)
	if pin != body.Pin {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid pin",
		})
	}
	tokenString, err := utility.GenerateJWTToken("doctor", body.License, 24*30*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create JWT Token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"token",
		tokenString,
		3600*24*30,
		"",
		"",
		false,
		true,
	)
	c.JSON(http.StatusAccepted, gin.H{
		"type":    "doctor",
		"license": body.License,
		"name":    name,
		"email":   email,
	})
}
