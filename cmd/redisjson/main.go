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

	//* example json
	user := struct{
		ID string
		Name string
	}{
		ID: "user-1",
		Name: "User 1",
	}

	valueSetResult, errSetResult := redishelper.JSONSet("test-go-json", "$", user)
	if errSetResult != nil {
		fmt.Printf("redishelper set error: %v\n", errSetResult)
	}
	fmt.Printf("redishelper set value: %v\n", valueSetResult)
	/**
		* when set cache is success
		redishelper set value: OK
	*/

	valueGetResult, errGetResult := redishelper.JSONGet("test-go-json")
	if errGetResult != nil {
		fmt.Printf("redishelper get error: %v\n", errGetResult)
	}
	fmt.Printf("redishelper get value: %v\n", valueGetResult)
	/**
		* when cache is exist
		redishelper get value: {"ID":"user-1","Name":"User 1"}
	*/

	/**
		* when cache is not exist
		redishelper get value: 
	*/
}