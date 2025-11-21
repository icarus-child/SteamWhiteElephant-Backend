package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/player", getPlayer)
	r.GET("/room-players", getRoomPlayers)
	r.POST("/player", createPlayer)

	r.GET("/room-exists", getRoomExists)
	r.GET("/retrieve-all", getAll)

	r.GET("/room-presents", getRoomGifts)
	r.POST("/present", createGift)

	return r
}
