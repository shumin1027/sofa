package connect

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"testing"
	"time"
)

var config = &RedisConfig{
	Address: []string{"localhost:6379"},
	DB:      0,
}

var rdb redis.UniversalClient

func init() {
	InitRedis(config)
	rdb = RedisClient()
}

func TestInitRedis(t *testing.T) {
	pong, _ := rdb.Ping().Result()
	fmt.Println(pong)
}

func TestList(t *testing.T) {
	go loop()
	for {
		time.Sleep(1000)
	}
}

func loop() {
	for {
		// timeout 5 s
		msg := rdb.BRPop(0*1000*1000*1000, "shumin")
		fmt.Println(msg.String())
	}
}

func TestPubSub(t *testing.T) {
	pubsub := rdb.Subscribe("shumin")

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	//Publish a message.
	err = rdb.Publish("shumin", "hello").Err()
	if err != nil {
		panic(err)
	}

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})

	// Consume messages.
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
