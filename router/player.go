package router

import (
	"log"

	"main/db"

	"github.com/gin-gonic/gin"
)

func getPlayer(ctx *gin.Context) {
	id := ctx.Query("id")

	player, err := db.GetPlayer(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"name":  nil,
			"room":  nil,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"name":   player.Name,
		"roomId": player.RoomID,
		"error":  nil,
	})
}

func getRoomPlayers(ctx *gin.Context) {
	id := ctx.Query("id")

	players, err := db.GetRoomPlayers(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"error":   err.Error(),
			"players": nil,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"error":   nil,
		"players": players,
	})
}

func createPlayer(ctx *gin.Context) {
	var json db.Player
	err := ctx.BindJSON(&json)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	player := db.Player{
		Name: json.Name,
	}
	err = db.CreatePlayer(ctx, json.PlayerID, json.RoomID, player)
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
	return
}
