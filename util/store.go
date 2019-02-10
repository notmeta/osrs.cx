package util

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis"
	"log"
)

var Store *redis.Client

const RsnKeyFormat = "rsn:%s"

func StoreInit() {
	Store = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := Store.Ping().Result()
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
