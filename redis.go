package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var REDIS_HOST = "REDIS_HOST"
var REDIS_PORT = "REDIS_PORT"
var REDIS_PASSWORD = "REDIS_PASSWORD"

func pollRedis() (bool, error) {
	host := getEnvVar(REDIS_HOST, "")
	if host == "" {
		return false, fmt.Errorf("REDIS_HOST is not set")
	}

	port := getEnvVar(REDIS_PORT, "6379")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return false, fmt.Errorf("REDIS_PORT is not a valid integer: %w", err)
	}

	password := getEnvVar(REDIS_PASSWORD, "")

	rdb := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", host, portInt),
		Password:    password,
		DB:          0,
		DialTimeout: 2 * time.Second,
		ReadTimeout: 2 * time.Second,
	})
	defer rdb.Close()

	_, err = rdb.Ping(context.TODO()).Result()
	if err != nil {
		return true, err
	}

	fmt.Printf("Redis is available at %s:%d\n", host, portInt)
	return false, nil
}
