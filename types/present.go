package types

type Gift struct {
	GifterID string   `json:"gifter"`
	SteamID  int      `json:"steamId"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`
}

type ItemJson struct {
	Name    string   `json:"name"`
	SteamId int      `json:"gameId"`
	Tags    []string `json:"tags"`
}

type PresentJson struct {
	GifterId string     `json:"gifterId"`
	Items    []ItemJson `json:"items"`
}
