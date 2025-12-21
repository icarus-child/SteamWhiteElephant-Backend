package router

import (
	"log"

	"main/db"

	"github.com/gin-gonic/gin"
)

func getRoomExists(ctx *gin.Context) {
	id := ctx.Query("id")

	exists, err := db.CheckRoomExists(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"exists": nil,
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"exists": exists,
		"error":  nil,
	})
}

func getRoomStarted(ctx *gin.Context) {
	id := ctx.Query("id")

	room, err := db.GetRoom(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"started": nil,
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"started": room.Started,
		"error":   nil,
	})
}

func startRoom(ctx *gin.Context) {
	id := ctx.Query("id")
	var room db.NarrowRoom
	room.RoomID = id
	room.Started = true

	err := db.CreateRoom(ctx, room)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"error": nil,
	})
}
