package controllers

import (
	"documenter/models/requests"
	"documenter/services"

	"github.com/gin-gonic/gin"
)

func GenerateDocController(c *gin.Context) {
	// Set headers for SSE
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	var request requests.GenDocRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.SSEvent("error", gin.H{"error": err.Error()})
		return
	}

	// Create a channel for status updates
	statusChan := make(chan string)

	// Call service asynchronously
	go func() {
		response, err := services.GenerateDocService(request, statusChan)
		if err != nil {
			c.SSEvent("error", gin.H{"error": err.Error()})
			close(statusChan)
			return
		}
		c.SSEvent("complete", gin.H{"doc": response})
		close(statusChan)
	}()

	// Send status updates to client
	for status := range statusChan {
		c.SSEvent("status", gin.H{"message": status})
		c.Writer.Flush()
	}
}
