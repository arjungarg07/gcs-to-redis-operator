package utils

import (
	"bufio"
	"cloud.google.com/go/storage"
	"context"
	"gcs-redis-operator/dto"
	protos "github.com/ShareChat/sharechat-central-protos/proto/pinternal/sc/feed_relevance_service"
	"github.com/bytedance/sonic"
	"github.com/golang/protobuf/proto"
	"log"
	"strconv"
	"time"
)

func downloadJsonFromGCS(ctx context.Context, bucketName string, objectPath string, dataChannel chan dto.PostStaticRedisMsetMsg, done chan bool) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create storage client: %v", err)
		return err
	}
	defer client.Close()

	rc, err := client.Bucket(bucketName).Object(objectPath).NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to open GCS object: %v", err)
		return err
	}
	defer rc.Close()
	scanner := bufio.NewScanner(rc)
	recordCount := 0
	for scanner.Scan() {
		recordCount++
		var post dto.PostStaticFeatures
		if err := sonic.Unmarshal(scanner.Bytes(), &post); err != nil {
			log.Printf("Failed to unmarshal JSON data: %v", err)
			continue // Skip this line and move to the next one
		}

		createdOnInt, _ := strconv.ParseInt(post.CreatedOn, 10, 64)
		contentTypeflt64, _ := strconv.ParseFloat(post.ContentType, 32)
		contentType := float32(contentTypeflt64)
		duration := float32(post.Duration)
		predictedProb := float32(post.PredictedProb)
		commentOff, _ := strconv.ParseInt(post.CommentOff, 10, 64)
		postShareDisabled, _ := strconv.ParseInt(post.PostShareDisabled, 10, 64)
		height, _ := strconv.ParseInt(post.Height, 10, 64)
		width, _ := strconv.ParseInt(post.Width, 10, 64)
		// Process the single PostStaticFeatures struct
		postFeature := &protos.PostStaticFeatures{
			PostId:            post.PostID,
			TagId:             post.TagID,
			CreatorId:         post.CreatorID,
			Language:          post.Language,
			CreatedOn:         createdOnInt,
			LOTopic:           post.LOTopic,
			L1Topic:           post.L1Topic,
			L2Topic:           post.L2Topic,
			CU_L1Topic:        post.CUL1Topic,
			Duration:          duration,
			CreatorIp:         post.CreatorIP,
			CreatorCity:       post.CreatorCity,
			CreatorState:      post.CreatorState,
			CreatorGender:     post.CreatorGender,
			CreatorBadge:      post.CreatorBadge,
			CreatorType:       post.CreatorType,
			PredictedProb:     predictedProb,
			PredictedTopic:    post.PredictedTopic,
			ContentType:       contentType,
			TagGenre:          post.TagGenre,
			Badge:             post.Badge,
			L0Taxonomy:        post.L0Taxonomy,
			L1Taxonomy:        post.L1Taxonomy,
			L2Taxonomy:        post.L2Taxonomy,
			L3Taxonomy:        post.L3Taxonomy,
			L4Taxonomy:        post.L4Taxonomy,
			CommentOff:        commentOff,
			PostShareDisabled: postShareDisabled,
			Height:            height,
			Width:             width,
			HybridTopic:       post.HybridTopic,
		}
		redisValue, err := proto.Marshal(postFeature)
		if err != nil {
			log.Printf("Error while scanning JSON data")
			return err
		}
		redisMsg := dto.PostStaticRedisMsetMsg{
			Key: post.PostID,
			Val: redisValue,
		}
		dataChannel <- redisMsg
	}
	log.Printf("Total records able to read from GCS bucket having object paths as %s are %v", objectPath, recordCount)
	time.Sleep(15 * time.Second)
	done <- true
	close(dataChannel)
	if err := scanner.Err(); err != nil {
		log.Printf("Error while scanning JSON data")
		return err
	}
	return err
}

func ReadGcsUncompressedFile(ctx context.Context, bucketName string, objectPath string, dataChannel chan dto.PostStaticRedisMsetMsg, done chan bool) {
	err := downloadJsonFromGCS(ctx, bucketName, objectPath, dataChannel, done)
	if err != nil {
		log.Fatal(err)
	}
}
