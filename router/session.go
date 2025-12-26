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

	err := db.StartRoom(ctx, id)
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

func getRoomTurnIndex(ctx *gin.Context) {
	id := ctx.Query("id")
	index, err := db.GetRoomTurnIndex(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"index": 0,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"index": index,
		"error": nil,
	})
}

type IndexRequest struct {
	Index int16 `json:"index"`
}

func setRoomTurnIndex(ctx *gin.Context) {
	var req IndexRequest
	id := ctx.Query("id")
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = db.SetRoomTurnIndex(ctx, id, req.Index)
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

func randomizePlayerOrder(ctx *gin.Context) {
	id := ctx.Query("id")
	err := db.RandomizePlayerOrder(ctx, id)
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
