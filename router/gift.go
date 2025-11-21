package router

import (
	"log"
	"main/db"
	"main/types"
	"main/utility"

	"github.com/gin-gonic/gin"
)

func createGift(ctx *gin.Context) {
	var present types.PresentJson
	err := ctx.BindJSON(&present)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, item := range present.Items {
		var gift types.Gift = types.Gift{
			GifterID: present.GifterId,
			SteamID:  item.SteamId,
			Name:     item.Name,
			Tags:     item.Tags,
		}
		err = db.CreateGift(ctx, gift)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	ctx.JSON(200, gin.H{
		"error": nil,
	})
	return
}

func getRoomGifts(ctx *gin.Context) {
	id := ctx.Query("id")

	gifts, err := db.GetRoomGifts(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"presents": nil,
			"error":    err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"presents": utility.CollectPresentsByGifter(gifts),
		"error":    nil,
	})
}
