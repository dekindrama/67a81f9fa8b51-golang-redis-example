package redishelper

import (
	"example/golang-redis-example/datasources/redis"
	"example/golang-redis-example/helpers/timehelper/v1"
)

func Set(key, value string, seconds int) (string, error) {
	resultValue, resultErr := redis.Client.Set(
			redis.Ctx,
			key,
			value,
			timehelper.Seconds(seconds),
		).Result()
	return resultValue, resultErr
}

func Get(key string) (string, error) {
	resultValue, resultErr := redis.Client.Get(
		redis.Ctx,
		key,
	).Result()
	return resultValue, resultErr
}

func JSONSet(key, path string, value interface{}) (string, error) {
	resultValue, resultErr := redis.Client.JSONSet(
			redis.Ctx,
			key,
			path,
			value,
		).Result()
	return resultValue, resultErr
}

func JSONGet(key string) (string, error) {
	resultValue, resultErr := redis.Client.JSONGet(
		redis.Ctx,
		key,
	).Result()
	return resultValue, resultErr;
}
