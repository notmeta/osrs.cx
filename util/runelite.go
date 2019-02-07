package util

import (
	"encoding/json"
	"fmt"
)

const runeliteBootstrapUrl = "https://static.runelite.net/bootstrap.json"
const runeliteApiUrl = "https://api.runelite.net/runelite-%s/"

type bootstrap struct {
	Client client
}

type client struct {
	Version string
}

// TODO cache this (and version) to prevent unnecessary requests
func RuneliteApiUrl() (url string) {
	return fmt.Sprintf(runeliteApiUrl, LatestRuneliteVersion())
}

func LatestRuneliteVersion() (version string) {
	url := runeliteBootstrapUrl
	body, _ := GetBody(&url)

	bs := bootstrap{}
	_ = json.Unmarshal(body, &bs)

	return bs.Client.Version
}
