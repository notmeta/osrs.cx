package util

import (
	"encoding/json"
	"fmt"
	"github.com/sahilm/fuzzy"
	"github.com/tevino/abool"
	"log"
	"sync"
	"time"
)

type itemIdName struct {
	Id   int
	Name string
}

type itemDict []itemIdName

func (items itemDict) String(i int) string {
	return items[i].Name
}

func (items itemDict) Len() int {
	return len(items)
}

const itemCacheKey = "itemcache"

var items itemDict
var rwm sync.RWMutex
var updatingCache = abool.New()

func FindBestMatchItemId(input string) *int {
	results := FindItemIds(input)

	if results == nil {
		return nil
	}

	return &results[0]
}

func FindItemIds(input string) []int {
	updateItemCache()

	rwm.RLock()
	defer rwm.RUnlock()

	// find best match based on string input
	results := fuzzy.FindFrom(input, items)

	if results.Len() == 0 {
		return nil
	}

	toReturn := make([]int, results.Len())

	for idx, r := range results {
		toReturn[idx] = items[r.Index].Id
	}

	return toReturn
}

func updateItemCache() {
	_, err := Store.Get(itemCacheKey).Result()

	if items.Len() != 0 && err == nil {
		return
	}

	if !updatingCache.SetToIf(false, true) {
		return
	}

	log.Println("Local item cache expired/not found - populating...")

	rwm.Lock()
	defer rwm.Unlock()
	defer updatingCache.UnSet()
	defer Store.Set(itemCacheKey, 1, 6*time.Hour)

	items = itemDict{} // clear the slice of cached items

	url := fmt.Sprintf("%sitem/prices", RuneliteApiUrl())
	body, _ := GetBody(&url)

	err = json.Unmarshal(body, &items)

	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Local item cache populated with %d entries\n", items.Len())
}
