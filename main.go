package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	sentinelAddrs = []string{"redis-sentinel-1:26379", "redis-sentinel-2:26379", "redis-sentinel-3:26379"}
	masterName    = "mymaster"
	channel       = "mychannel"
	ctx           = context.Background()
)

func main() {
	rdb := connectToRedis()
	go func() {
		for {
			publishMessage(rdb)
			time.Sleep(1 * time.Second)
		}
	}()
	subscribeMessage(rdb)
}

func connectToRedis() *redis.Client {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    masterName,
		SentinelAddrs: sentinelAddrs,
        DialTimeout:   1 * time.Second,
        ReadTimeout:   1 * time.Second,
	})
	return rdb
}

func publishMessage(rdb *redis.Client) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	err := rdb.Publish(ctx, channel, currentTime).Err()
	if err != nil {
		log.Printf("Publish error: %v", err)
	}
}

func subscribeMessage(rdb *redis.Client) {
	sub := rdb.Subscribe(ctx, channel)
	defer sub.Close()

	ch := sub.Channel()

	for msg := range ch {
		fmt.Printf("Received message from %s: %s\n", msg.Channel, msg.Payload)
	}
}
