package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test connection
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err) // should print: PONG <nil>

	// Set a key
	err = rdb.Set(ctx, "greeting", "hello redis", 0).Err()
	if err != nil {
		panic(err)
	}

	// Get the key
	val, err := rdb.Get(ctx, "greeting").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("greeting:", val)

	// Nonexistent key
	val2, err := rdb.Get(ctx, "unknown").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("unknown:", val2)
	}
}
