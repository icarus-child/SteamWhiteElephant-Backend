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
