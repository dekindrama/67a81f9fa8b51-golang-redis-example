package main

import (
	"example/golang-redis-example/config"
	"example/golang-redis-example/datasources/redis"
	"example/golang-redis-example/helpers/redishelper/v1"
	"fmt"
	"log"
)

func main() {
	//* connect redis
	errConnectRedis := redis.ConnectRedis(
		config.Config("REDIS_ADDRESS"),
		config.Config("REDIS_PASSWORD"),
	)
	if errConnectRedis != nil {
		log.Fatalf("redis error main: %v\n", errConnectRedis)
	}
	fmt.Printf("redis Client: %v\n", redis.Client)
	fmt.Printf("redis Ctx: %v\n", redis.Ctx)

	valueSetResult, errSetResult := redishelper.Set("test-go", "hello world", 90)
	if errSetResult != nil {
		fmt.Printf("redishelper set error: %v\n", errSetResult)
	}
	fmt.Printf("redishelper set value: %v\n", valueSetResult)
	/**
		* when set cache is success
		redishelper set value: OK
	*/


	valueGetResult, errGetResult := redishelper.Get("test-go")
	if errGetResult != nil {
		fmt.Printf("redishelper get error: %v\n", errGetResult)
	}
	fmt.Printf("redishelper get value: %v\n", valueGetResult)
	/**
		* when cache is exist
		redishelper get value: hello world
	*/

	/**
		* when cache is not exist
		redishelper get error: redis: nil
		redishelper get value: 
	*/
}