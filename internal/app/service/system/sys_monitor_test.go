package system

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestMonitor(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	// val, _ := client.Get(ctx, "key").Result()
	// fmt.Println(val)
	fmt.Println(client.DBSize(ctx))
	fmt.Println(client.Info(ctx, "Clients").Val())
	// fmt.Printf("%T", client.Info(ctx, "Stats").Val())
}
