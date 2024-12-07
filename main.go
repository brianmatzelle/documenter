package main

import (
	"documenter/config"
	c "documenter/controllers"
	"fmt"
	"log"
)

func main() {
	config.LoadEnv()
	router := c.SetupRouter()
	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting http server on port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
