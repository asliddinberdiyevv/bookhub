package db

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

func ConnectRS() *redis.Client {
	rs_host := os.Getenv("REDIS_HOST")
	rs_port := os.Getenv("REDIS_PORT")
	rs_password := os.Getenv("REDIS_PASSWORD")
	rs_Addr := fmt.Sprintf("%s:%s", rs_host, rs_port)

	client := redis.NewClient(&redis.Options{
		Addr:     rs_Addr,
		Password: rs_password,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	logrus.Infof("redis ping status: %s", pong)
	if err != nil {
		logrus.Fatalf("failed connecting to redis db error: %s", err.Error())
		os.Exit(1)
	}
	logrus.Infof("successfully connected to redis db")

	return client
}
