package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 16 << 20

	r.GET("/player", getPlayer)
	r.GET("/room-players", getRoomPlayers)
	r.POST("/player", createPlayer)

	r.GET("/room-exists", getRoomExists)

	r.GET("/room-presents", getRoomGifts)
	r.GET("/texture", getTexture)
	r.POST("/present", createGift)
	r.POST("/texture", createTexture)

	r.GET("/room-started", getRoomStarted)
	r.POST("/start-room", startRoom)

	return r
}
