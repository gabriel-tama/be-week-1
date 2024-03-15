package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gabriel-tama/be-week-1/types/response"
	"github.com/gin-gonic/gin"
)

type ImageController struct {
	jwtService services.JWTService
	s3Service  services.S3Service
}

func NewImageController(jwtService services.JWTService, s3Service services.S3Service) *ImageController {
	return &ImageController{jwtService: jwtService, s3Service: s3Service}
}

func (ic *ImageController) UploadImage(c *gin.Context) {
	var res = response.ImageResponse{}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "file not found"})
		return
	}

	// Check file size
	if file.Size > 2*1024*1024 || file.Size < 10*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "image is wrong (not *.jpg | *.jpeg, more than 2MB or less than 10KB)"})
		return
	}

	fileHeader, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	defer fileHeader.Close()

	buffer := make([]byte, 512) // Why 512 bytes? See http://golang.org/pkg/net/http/#DetectContentType
	_, err = fileHeader.Read(buffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	contentType := http.DetectContentType(buffer)
	if contentType != "image/jpeg" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "image is wrong (not *.jpg | *.jpeg, more than 2MB or less than 10KB)"})
		return
	}

	fileBuffer, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "file not found"})
		return
	}

	objKey := fmt.Sprintf("%s/%v-%s", "ngab-gab", time.Now().Unix(), file.Filename)

	_, err = ic.s3Service.UploadFile(objKey, fileBuffer, contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}

	res.Url = ic.s3Service.GetObjectWithUrl(objKey)

	c.JSON(http.StatusOK, gin.H{"message": "image uploaded successfully", "data": res})
}
