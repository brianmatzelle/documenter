package main

import (
	"documenter/config"
	c "documenter/controllers"
	"fmt"
)

func main() {
	config.LoadEnv()
	router := c.SetupRouter()
	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}
