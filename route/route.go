package route

import (
	. "gin/ginDemo/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/person", AddPerson)

	router.GET("/persons", GetPersonsApi)

	router.GET("/person/:id", GetPersonApi)

	router.PUT("/person/:id", ModPerson)

	router.DELETE("/person/:id", DelPersonApi)

	return router
}