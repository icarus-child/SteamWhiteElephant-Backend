package router

import (
	"log"
	"main/db"

	"github.com/gin-gonic/gin"
)

type itemJson struct {
	Name    string   `json:"name"`
	SteamId int      `json:"gameId"`
	Tags    []string `json:"tags"`
}

type presentJson struct {
	GifterId string     `json:"gifterId"`
	Items    []itemJson `json:"items"`
}

func createGift(ctx *gin.Context) {
	var present presentJson
	err := ctx.BindJSON(&present)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, item := range present.Items {
		var gift db.Gift = db.Gift{
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

func collectPresentsByGifter(gifts []db.Gift) (presents []presentJson) {
	if len(gifts) == 0 {
		return
	}

	var itemsMap map[string][]itemJson = make(map[string][]itemJson)

	for _, gift := range gifts {
		if itemsMap[gift.GifterID] == nil {
			itemsMap[gift.GifterID] = make([]itemJson, 0)
		}
		itemsMap[gift.GifterID] = append(itemsMap[gift.GifterID], itemJson{
			Name:    gift.Name,
			SteamId: gift.SteamID,
			Tags:    gift.Tags,
		})
	}

	for gifterId, items := range itemsMap {
		var present presentJson = presentJson{
			GifterId: gifterId,
			Items:    items,
		}
		presents = append(presents, present)
	}
	return
}

func getRoomGifts(ctx *gin.Context) {
	id := ctx.Query("id")

	gifts, err := db.GetRoomGifts(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, gin.H{
			"gifts": nil,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"gifts": collectPresentsByGifter(gifts),
		"error": nil,
	})
}
