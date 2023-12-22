package redis

import (
	"context"
	"fmt"
	"gcs-redis-operator/config"
	"gcs-redis-operator/constants"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type RedisClient struct {
	RedisClusterClient *redis.ClusterClient
	ErrorRate          float64
	ExpiryTimeout      time.Duration
}

func NewRedisClusterClient(config config.RedisConfig, expiry string) RedisClient {
	redisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:      []string{fmt.Sprintf("%s:%d", config.Host, config.Port)},
		Password:   config.Password,
		MaxRetries: 3,
	})
	errorRate := float64(constants.REDIS_ERROR_RATE)
	expiryInSeconds, err := strconv.Atoi(expiry)
	fmt.Printf("Expiry for redis keys to be set is: %v seconds \n", expiryInSeconds)
	if err != nil {
		log.Fatalf("error in converting expiry to integer %v", err.Error())
	}

	ctx := context.Background()
	var pingCounter int
	for pingCounter = 0; pingCounter < 5; pingCounter++ {
		_, err = redisClient.Ping(ctx).Result()
		if err != nil {
			log.Printf("error in creating redis connection %v", err.Error())
		} else {
			break
		}
		time.Sleep(time.Duration(rand.Intn(5)+10) * time.Second)
	}
	if pingCounter == 5 {
		log.Fatalf("error in creating redis connection %v", err.Error())
	}
	expiryTimeout := time.Duration(expiryInSeconds) * time.Second
	return RedisClient{redisClient, errorRate, expiryTimeout}
}
