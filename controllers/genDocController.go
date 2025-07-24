package controllers

import (
	"documenter/config"
	"documenter/models/requests"
	"documenter/services"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GenerateDocController(c *gin.Context) {
	// Set headers for SSE
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	// Get GitLab token from environment to validate it exists
	gitlabToken := config.GetEnv("GITLAB_TOKEN")
	if gitlabToken == "" {
		c.SSEvent("error", gin.H{"error": "GITLAB_TOKEN environment variable is not set"})
		return
	}

	// Parse query parameters
	mrLinksJSON := c.Query("mrLinks")
	var mrLinks []string
	if err := json.Unmarshal([]byte(mrLinksJSON), &mrLinks); err != nil {
		c.SSEvent("error", gin.H{"error": "Invalid mrLinks format"})
		return
	}

	request := requests.GenDocRequest{
		MrLinks: mrLinks,
		Model:   c.Query("model"),
	}

	// Validate the request
	if len(request.MrLinks) == 0 {
		c.SSEvent("error", gin.H{"error": "mrLinks is required"})
		return
	}
	if request.Model == "" {
		c.SSEvent("error", gin.H{"error": "model is required"})
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

func GenFromAuthorController(c *gin.Context) {
	// Set headers for SSE
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	// Get GitLab token from environment to validate it exists
	gitlabToken := config.GetEnv("GITLAB_TOKEN")
	if gitlabToken == "" {
		c.SSEvent("error", gin.H{"error": "GITLAB_TOKEN environment variable is not set"})
		return
	}

	author := c.Query("author")
	model := c.Query("model")

	statusChan := make(chan string)
	// Call service asynchronously
	go func() {
		response, err := services.GenDocFromAuthor(author, model, statusChan)
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
