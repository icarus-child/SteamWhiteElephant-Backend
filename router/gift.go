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

	err = db.CreateGiftName(ctx, present.GifterId, present.GiftName)
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
	ctx.JSON(200, gin.H{"ok": true})
}

func getPlayerHeldGift(ctx *gin.Context) {
	id := ctx.Query("id")

	gifts, err := db.GetPlayerHeldGifts(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"present": nil,
			"error":   err.Error(),
		})
		return
	}

	if len(gifts) == 0 {
		ctx.JSON(200, gin.H{
			"present": nil,
			"error":   nil,
		})
		return
	}

	presents := utility.CollectPresentsByGifter(ctx, gifts)
	ctx.JSON(200, gin.H{
		"present": presents[0],
		"error":   nil,
	})
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

func getPresentStolenThisRound(ctx *gin.Context) {
	id := ctx.Query("id")

	opened, err := db.GetPresentStolenThisRound(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"opened": false,
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"opened": opened,
		"error":  nil,
	})
}

func resetRound(ctx *gin.Context) {
	id := ctx.Query("id")
	err := db.ResetRound(ctx, id)
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

func markPresentStolen(ctx *gin.Context) {
	id := ctx.Query("id")
	err := db.MarkPresentStolen(ctx, id)
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

type TakeRequest struct {
	PlayerID  string `json:"playerId"`
	PresentID string `json:"presentId"`
}

func takeOrStealPresent(ctx *gin.Context) {
	var req TakeRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = db.TakeOrStealPresent(ctx, req.PlayerID, req.PresentID)
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

func increaseTimesStolen(ctx *gin.Context) {
	id := ctx.Query("id")
	err := db.IncreaseTimesStolen(ctx, id)
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

func getTimesStolen(ctx *gin.Context) {
	id := ctx.Query("id")
	timesStolen, err := db.GetTimesStolen(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"timesStolen": nil,
			"error":       err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"timesStolen": timesStolen,
		"error":       nil,
	})
}
