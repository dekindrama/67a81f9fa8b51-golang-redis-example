package main

import (
	"example/golang-redis-example/config"
	"example/golang-redis-example/datasources/redis"
	"example/golang-redis-example/helpers/redishelper/v1"
	"fmt"
	"log"
)

func main() {
	//* init redis client
	errConnectRedis := redis.ConnectRedis(
		config.Config("REDIS_ADDRESS"), //* localhost:6379
		config.Config("REDIS_PASSWORD"), //* secret
	)
	if errConnectRedis != nil {
		log.Fatalf("redis error main: %v\n", errConnectRedis)
	}
	fmt.Printf("redis Client: %v\n", redis.Client)
	fmt.Printf("redis Ctx: %v\n", redis.Ctx)

	//* set redis
	valueSetResult, errSetResult := redishelper.Set(
		"test-go", //* key
		"hello world", //* value
		90, //* data duration, 90s
	)
	if errSetResult != nil {
		fmt.Printf("redishelper set error: %v\n", errSetResult)
	}
	fmt.Printf("redishelper set value: %v\n", valueSetResult)
	/**
		* when set cache is success
		redishelper set value: OK
	*/


	//* get redis
	valueGetResult, errGetResult := redishelper.Get(
		"test-go", //* key
	)
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

	//* Delete redis
	valueDelResult, errDelResult := redishelper.Del(
		"test-go", //* key
	)
	if errDelResult != nil {
		fmt.Printf("redishelper Del error: %v\n", errDelResult)
	}
	fmt.Printf("redishelper Del value: %v\n", valueDelResult)
	/**
		* when del cache is success
		redishelper Del value: 1
	*/
}