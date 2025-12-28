package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 16 << 20

	r.GET("/player", getPlayer)
	r.GET("/room-players", getRoomPlayers)
	r.GET("/room-ordered-players", getOrderedRoomPlayers)
	r.POST("/player", createPlayer)

	r.GET("/room-exists", getRoomExists)

	r.GET("/room-presents", getRoomGifts)
	r.GET("/texture", getTexture)
	r.GET("/player-present", getPlayerHeldGift)
	r.GET("/present-stolen-this-round", getPresentStolenThisRound)
	r.GET("/times-stolen", getTimesStolen)
	r.POST("/present", createGift)
	r.POST("/texture", createTexture)
	r.POST("/mark-present-stolen", markPresentStolen)
	r.POST("/take-or-steal-present", takeOrStealPresent)
	r.POST("/increase-times-stolen", increaseTimesStolen)

	r.GET("/room-started", getRoomStarted)
	r.GET("/room-turn-index", getRoomTurnIndex)
	r.POST("/start-room", startRoom)
	r.POST("/reset-round", resetRound)
	r.POST("/room-turn-index", setRoomTurnIndex)
	r.POST("/randomize-player-order", randomizePlayerOrder)

	return r
}
