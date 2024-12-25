package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/player", getPlayer)
	r.POST("/player", createPlayer)

	return r
}
