package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ItemSearchResult struct {
	Items []Item
}

type Item struct {
	Id          int
	Name        string
	Description string
	Type        string
}

func SearchItem(query string) ItemSearchResult {
	query = strings.Replace(query, " ", "+", -1)
	url := fmt.Sprintf("%sitem/search?query=%s", RuneliteApiUrl(), query)

	body, _ := GetBody(&url)

	result := ItemSearchResult{}
	_ = json.Unmarshal(body, &result)

	return result
}

func (item *Item) GetIconUrl() string {
	return fmt.Sprintf("%sitem/%d/icon", RuneliteApiUrl(), item.Id)
}
