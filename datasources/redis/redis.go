package redis

import (
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"
)

var Client *goredis.Client
var Ctx context.Context

func ConnectRedis(addr, password string) error {
	ctx := context.Background()
	client := goredis.NewClient(&goredis.Options{
        Addr:	  addr,
        Password: password,
        DB:		  0,  // Use default DB
        Protocol: 2,  // Connection protocol
    })

	//* test ping
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Printf("redis ping: %v\n", pong)

	//* assign to global var
	Client = client
	Ctx = ctx

	return nil
}



