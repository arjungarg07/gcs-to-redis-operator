package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"log"
	"os"
)

type RedisConfig struct {
	Host     string `json:"host"  validate:"required"`
	Port     int    `json:"port"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type Configuration struct {
	PostStaticRedis RedisConfig `json:"post_static_redis" validate:"required"`
}

var Config *Configuration

func IsStagingEnv() bool {
	env := os.Getenv("ACTIVE_ENV")
	if env == "STAGING" {
		return true
	}
	return false
}

func IsProdEnv() bool {
	env := os.Getenv("ACTIVE_ENV")
	if env == "PRODUCTION" {
		return true
	}
	return false
}

func InitConfig() {
	Config = new(Configuration)
	dirname, dirErr := os.Getwd()
	if dirErr != nil {
		log.Fatalf("error in getting directory, err=%+v", dirErr)
	}

	var fileName string
	if IsProdEnv() {
		fileName = dirname + "/config/config.prod.json"
	} else if IsStagingEnv() {
		fileName = dirname + "/config/config.stage.json"
	} else {
		fileName = dirname + "/config/config.dev.json"
	}
	err := gonfig.GetConf(fileName, Config)
	if err != nil {
		panic(fmt.Sprintf("could not initialize config, err=%+v\n", err))
	}
}

func GetRedisConfig(redisName string) RedisConfig {
	if redisName == "postStaticRedis" {
		return Config.PostStaticRedis
	} else {
		panic("No matching Redis")
	}
}
