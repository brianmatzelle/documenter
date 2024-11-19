package main

import (
	"documenter/config"
	c "documenter/controllers"
	"fmt"
)

func main() {
	router := c.SetupRouter()
	router.Run(fmt.Sprintf(":%s", config.PORT))
}
