package main

import (
	"documenter/config"
	c "documenter/controllers"
	"fmt"
)

func main() {
	config.LoadEnv()
	router := c.SetupRouter()
	router.Run(fmt.Sprintf(":%s", config.GetEnv("PORT")))
}
