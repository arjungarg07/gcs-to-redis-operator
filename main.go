package main

import (
	"context"
	"flag"
	"fmt"
	"gcs-redis-operator/config"
	"gcs-redis-operator/constants"
	"gcs-redis-operator/dto"
	"gcs-redis-operator/redis"
	"gcs-redis-operator/utils"
	"github.com/bytedance/sonic"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	var (
		gcsPath        string
		redisInstance  string
		redisOperation string
		expiry         string
		batchSize      string
	)

	config.InitConfig()
	flag.StringVar(&gcsPath, "gcsPath", "", "Gcs Path")
	flag.StringVar(&redisInstance, "redisInstance", "", "Redis Instance ID")
	flag.StringVar(&redisOperation, "operation", "", "Redis operation to be performed")
	flag.StringVar(&expiry, "expiry", "", "Redis keys expiry in seconds")
	flag.StringVar(&batchSize, "batchSize", constants.DEFAULT_BATCH_SIZE, "Batch size to be pushed in redis from gcs")

	flag.Parse()

	if gcsPath == "" || redisInstance == "" || redisOperation == "" || expiry == "" {
		fmt.Println("Error: Required flags are missing or invalid.")
		flag.PrintDefaults()
		return
	}
	redisBatchSize, err := strconv.Atoi(batchSize)
	if err != nil {
		log.Fatal(err)
		return
	}
	var gcsBucketName, gcsObjectPath string
	gcsPathParts := strings.Split(gcsPath, "//")
	if len(gcsPathParts) < 2 {
		fmt.Println("Invalid gcs path")
		return
	}
	gcsBucketAndObjectPath := gcsPathParts[1]
	parts := strings.SplitN(gcsBucketAndObjectPath, "/", 2)

	if len(parts) == 2 {
		gcsBucketName = parts[0]
		gcsObjectPath = parts[1]
	} else {
		fmt.Println("Invalid gcs path")
		return
	}

	fmt.Println("Starting GCS to Redis job....")
	fmt.Println("Gcs Bucket Name:", gcsBucketName)
	fmt.Println("Gcs Object Path:", gcsObjectPath)
	fmt.Println("Redis Instance Id:", redisInstance)
	fmt.Println("Redis Operation to be performed:", redisOperation)
	fmt.Println("Batch Size:", redisBatchSize)

	redisConfig := config.GetRedisConfig(redisInstance)
	ctx := context.Background()
	redisClient := redis.NewRedisClusterClient(redisConfig, expiry)
	dataChannel := make(chan dto.PostStaticRedisMsetMsg, constants.DATA_CHANNEL_SIZE)
	done := make(chan bool, 1)
	var v dto.PostStaticFeatures
	err = sonic.Pretouch(reflect.TypeOf(v))
	if err != nil {
		log.Fatal(err)
	}
	go utils.ReadGcsUncompressedFile(ctx, gcsBucketName, gcsObjectPath, dataChannel, done)
	if redisOperation == "mset" {
		redisClient.ProcessMset(dataChannel, done, redisBatchSize)
	}
	os.Exit(0)
}
