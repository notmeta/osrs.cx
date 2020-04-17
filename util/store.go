package util

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var Store *redis.Client
var RedisIp string

const RsnKeyFormat = "rsn:%s"

func StoreInit() {
	Store = redis.NewClient(&redis.Options{
		Addr:     RedisIp + ":6379",
		Password: "",
		DB:       0,
	})

	var err error
	pingAttempts := 0

	for pingAttempts < 5 {
		_, err = Store.Ping().Result()

		if err != nil {
			log.Printf("Ping failed, waiting and trying again")
			pingAttempts += 1
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		log.Fatalf("Failed to connect to local redis container!\n%s", err)
	} else {
		log.Println("Connected to local redis container...")
	}
}

func GetRsn(author *discordgo.User) (rsn string) {
	rsn, _ = Store.Get(fmt.Sprintf(RsnKeyFormat, author.ID)).Result()
	return
}
