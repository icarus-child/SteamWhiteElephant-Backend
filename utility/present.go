package utility

import (
	"context"

	"main/db"
	"main/types"
)

func CollectPresentsByGifter(ctx context.Context, gifts []types.Gift) (presents []types.PresentJson) {
	if len(gifts) == 0 {
		return presents
	}

	var itemsMap map[string][]types.ItemJson = make(map[string][]types.ItemJson)

	for _, gift := range gifts {
		if itemsMap[gift.GifterID] == nil {
			itemsMap[gift.GifterID] = make([]types.ItemJson, 0)
		}
		itemsMap[gift.GifterID] = append(itemsMap[gift.GifterID], types.ItemJson{
			Name:    gift.Name,
			SteamId: gift.SteamID,
			Tags:    gift.Tags,
		})
	}

	for gifterId, items := range itemsMap {
		texture, err := db.GetTexture(ctx, gifterId)
		if err != nil {
			println(err.Error())
			return presents
		}
		var present types.PresentJson = types.PresentJson{
			GifterId: gifterId,
			Items:    items,
			Texture:  texture,
		}
		presents = append(presents, present)
	}
	return presents
}
