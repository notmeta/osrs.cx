package util

import (
	"encoding/json"
	"fmt"
	"time"
)

type RunescapeGEItem struct {
	Item RunescapeGEItemDetails
}

type RunescapeGEItemDetails struct {
	Icon        string
	IconLarge   string `json:"icon_large"`
	Id          int
	Type        string
	TypeIcon    string `json:"typeIcon"`
	Name        string
	Description string
	Current     RunescapeGEItemCurrentPrice
	Today       RunescapeGEItemChange
	Members     bool
	Day30       RunescapeGEItemChange
	Day90       RunescapeGEItemChange
	Day180      RunescapeGEItemChange
}

type RunescapeGEItemCurrentPrice struct {
	Trend string
	Price int
}

type RunescapeGEItemChange struct {
	Trend string
	Price string
}

const runescapeItemKeyFormat = "runescapeitem:%d"
const runescapeGEApiFormat = "https://services.runescape.com/m=itemdb_oldschool/api/catalogue/detail.json?item=%d"

func redisKey(id int) string {
	return fmt.Sprintf(runescapeItemKeyFormat, id)
}

func RunescapeItemForId(id int) RunescapeGEItem {
	item := RunescapeGEItem{}

	cachedItem, _ := Store.Get(redisKey(id)).Result()

	var jsonItem []byte

	if len(cachedItem) > 0 {
		jsonItem = []byte(cachedItem)
	} else {
		url := fmt.Sprintf(runescapeGEApiFormat, id)
		jsonItem, _ = GetBody(&url)
		_ = Store.Set(redisKey(id), string(jsonItem), 12*time.Hour).Err()
	}

	_ = json.Unmarshal(jsonItem, &item)

	return item
}
