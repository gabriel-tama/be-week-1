package controllers

import (
	"net/http"

	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gabriel-tama/be-week-1/types"
	"github.com/gin-gonic/gin"
)

type ImageController struct {
	jwtService services.JWTService
	s3Service  services.S3Service
}

func NewImageController(jwtService services.JWTService) *ImageController {
	return &ImageController{jwtService: jwtService}
}

func (ic *ImageController) UploadImage(c *gin.Context) {
	var bank types.RegisterBankData
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required details are missing or invalid"})
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "file not found"})
		return
	}
	fileBuffer, err := file.Open()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	_, err = ic.s3Service.UploadFile("test", fileBuffer)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "account added sucessfully", "data": bank})
}
