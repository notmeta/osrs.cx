package util

import (
	"encoding/json"
	"fmt"
	"time"
)

type OsbPrice struct {
	ItemId         int `json:"item_id"`
	BuyAverage     int `json:"buy_average"`
	SellAverage    int `json:"sell_average"`
	OverallAverage int `json:"overall_average"`
}

const priceKeyFormat = "price:%d"

func GetOSBPrice(id int) OsbPrice {
	price := OsbPrice{}

	key := fmt.Sprintf(priceKeyFormat, id)
	cachedPrice, _ := Store.Get(key).Result()

	var jsonPrice []byte

	if len(cachedPrice) > 0 {
		jsonPrice = []byte(cachedPrice)
	} else {
		url := fmt.Sprintf("%s/osb/ge?itemId=%d", RuneliteApiUrl(), id)
		jsonPrice, _ = GetBody(&url)
		_ = Store.Set(key, string(jsonPrice), 20*time.Minute).Err()
	}

	_ = json.Unmarshal(jsonPrice, &price)

	return price
}
