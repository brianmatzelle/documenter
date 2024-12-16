package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", Ping)
	router.GET("/generate-doc", GenerateDocController)
	router.GET("/gen-from-author", GenFromAuthorController)

	return router
}
