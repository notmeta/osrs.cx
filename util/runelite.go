package util

import (
	"encoding/json"
	"fmt"
	"time"
)

const runeliteBootstrapUrl = "https://static.runelite.net/bootstrap.json"
const runeliteApiUrl = "https://api.runelite.net/runelite-%s/"
const runeliteSRNUrl = "https://static.runelite.net/"

const runeliteApiUrlKey = "runelite-api-url"

type bootstrap struct {
	Client client
}

type client struct {
	Version string
}

func RuneliteApiUrl() (url string) {
	url, _ = Store.Get(runeliteApiUrlKey).Result()

	if len(url) == 0 {
		url = fmt.Sprintf(runeliteApiUrl, LatestRuneliteVersion())
		_ = Store.Set(runeliteApiUrlKey, url, 3*time.Hour).Err()
	}

	return
}

func LatestRuneliteVersion() (version string) {
	url := runeliteBootstrapUrl
	body, _ := GetBody(&url)

	bs := bootstrap{}
	_ = json.Unmarshal(body, &bs)

	return bs.Client.Version
}
