package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/SohamGhugare/IHP/ipfs"
	"github.com/SohamGhugare/IHP/utility"
	"github.com/gin-gonic/gin"
)

func StoreFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while fetching": err.Error()})
		return
	}
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while reading": err.Error()})
		return
	}
	defer src.Close()

	// Create a destination file on the server
	fpath := "dummy/uploads/" + file.Filename
	dst, err := os.Create(fpath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while creating": err.Error()})
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	_, err = io.Copy(dst, src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while copying": err.Error()})
		return
	}

	cid := ipfs.StoreFile(fpath)
	ecid, err := utility.EncryptString("0123456789ABCDEF", cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while encrypting": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cid":           cid,
		"encrypted_cid": ecid,
	})
}

func GetFile(c *gin.Context) {
	var body struct {
		Cid string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to bind",
		})
		return
	}

	dcid, err := utility.DecryptString("0123456789ABCDEF", body.Cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while decrypting": err.Error()})
		return
	}
	file, err := ipfs.GetStoredFile(dcid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while fetching": err.Error()})
		return
	}

	// Set the appropriate headers for file download
	c.Header("Content-Disposition", "attachment; filename=downloaded.png")
	c.Header("Content-Type", "application/octet-stream")
	c.File(file.Name())

}
