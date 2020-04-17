package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

func Get(url *string) (resp *http.Response, err error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", *url, nil)
	req.Header.Add("User-Agent", fmt.Sprintf("osrs.cx/%s (+https://github.com/notmeta/osrs.cx)", runtime.Version()))

	return client.Do(req)
}

func GetBody(url *string) (body []byte, err error) {
	resp, _err := Get(url)

	if _err == nil {
		defer resp.Body.Close()
	}

	return ioutil.ReadAll(resp.Body)
}
