package router

import (
	"log"
	"main/db"

	"github.com/gin-gonic/gin"
)

type jsonPlayer struct {
	Id string
	db.Player
}

func getPlayer(ctx *gin.Context) {
	id := ctx.Query("id")

	player, err := db.GetPlayer(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"name":  nil,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"name":  player.Name,
		"error": nil,
	})
}

func createPlayer(ctx *gin.Context) {
	var json jsonPlayer
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
	err = db.CreatePlayer(ctx, json.Id, player)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(400, gin.H{
		"error": nil,
	})
	return
}
