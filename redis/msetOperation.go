package redis

import (
	"context"
	"fmt"
	"gcs-redis-operator/dto"
	"log"
	"time"
)

func (redisClient *RedisClient) ProcessMset(dataChannel chan dto.PostStaticRedisMsetMsg, done chan bool, batchSize int) {
	messageList := make([]dto.PostStaticRedisMsetMsg, 0)
	errorCounter := 0
	totalCounter := 0
	timeTaken := 0.0
	for {
		select {
		case msg := <-dataChannel:
			{
				totalCounter += 1
				messageList = append(messageList, msg)
				if len(messageList) == batchSize {
					errorCount, t := redisClient.msetPipeline(messageList)
					messageList = make([]dto.PostStaticRedisMsetMsg, 0)
					errorCounter += errorCount
					timeTaken += t
					if float64(errorCounter)/float64(totalCounter) > redisClient.ErrorRate {
						log.Fatalf("High Error Rate")
					}
					if totalCounter%1000 == 0 {
						fmt.Printf("Total Entries Parsed: %d ,Avg Time Taken:%f ms , Total Errors: %d \n", totalCounter, timeTaken/float64(totalCounter)*1000, errorCounter)
					}
				}
			}
		case <-done:
			{
				if len(messageList) > 0 {
					errorCount, t := redisClient.msetPipeline(messageList)
					errorCounter += errorCount
					timeTaken += t
					if float64(errorCounter)/float64(totalCounter) > redisClient.ErrorRate {
						log.Fatalf("High Error Rate")
					}
				}
				fmt.Printf("Total Entries Parsed: %d ,Avg Time Taken:%f ms , Total Errors: %d \n", totalCounter, timeTaken/float64(totalCounter)*1000, errorCounter)
				return
			}
		}
	}
}

func (redisClient *RedisClient) msetPipeline(messageList []dto.PostStaticRedisMsetMsg) (int, float64) {
	pipe := redisClient.RedisClusterClient.Pipeline()
	ctx := context.Background()
	for _, msg := range messageList {
		pipe.SetEx(ctx, msg.Key, msg.Val, redisClient.ExpiryTimeout)
	}
	t := time.Now()
	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Println("Error in processing pipeline")
		return len(messageList), time.Since(t).Seconds()
	}
	return 0, time.Since(t).Seconds()
}
