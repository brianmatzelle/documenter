package controllers

import (
	"documenter/models/requests"
	"documenter/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateDocController(c *gin.Context) {
	var request requests.GenDocRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := services.GenerateDocService(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
