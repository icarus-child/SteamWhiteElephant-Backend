package router

import (
	"io"
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
}

func createTexture(ctx *gin.Context) {
	playerID := ctx.GetHeader("X-Player-ID")
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = db.CreateTexture(ctx, playerID, data)
	log.Println("received bytes:", len(data))
	ctx.JSON(200, gin.H{"ok": true})
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

	presents := utility.CollectPresentsByGifter(ctx, gifts)
	log.Println(len(presents[0].Texture))
	ctx.JSON(200, gin.H{
		"presents": presents,
		"error":    nil,
	})
}

func getTexture(ctx *gin.Context) {
	id := ctx.Query("id")

	texture, err := db.GetTexture(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Header("Content-Type", "image/png")
	ctx.Header("Content-Disposition", "inline; filename=texture.png")
	ctx.Writer.Write(texture)
}
