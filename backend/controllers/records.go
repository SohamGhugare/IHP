package controllers

import (
	"net/http"

	"github.com/SohamGhugare/IHP/contract"
	"github.com/gin-gonic/gin"
)

func AddRecord(c *gin.Context) {
	var body struct {
		IHP int
		Cid string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to bind",
		})
		return
	}

	tx := contract.AddRecord(body.IHP, body.Cid)
	c.JSON(http.StatusCreated, gin.H{
		"transaction_hash": tx,
	})
}
