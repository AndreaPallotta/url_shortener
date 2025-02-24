package store

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type StorageService struct {
	redisClient *redis.Client
}

type Url struct {
	ShortUrl string `json:"shortUrl"`
	FullUrl  string `json:"fullUrl"`
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong msg - {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMap(shortUrl string, ogUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, userId+":"+shortUrl, ogUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving url map | shortUrl: %s - ogUrl: %s - Err: %v", shortUrl, ogUrl, err))
	}

	err = storeService.redisClient.SAdd(ctx, "user:"+userId+":urls", shortUrl).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed adding shortUrl to user URLs set | userId: %s - shortUrl: %s - Err: %v", userId, shortUrl, err))
	}
}

func GetFullUrl(shortUrl string, userId string) string {
	res, err := storeService.redisClient.Get(ctx, userId+":"+shortUrl).Result()

	if err != nil {
		panic(fmt.Sprintf("Failed retrieving Url | shortUrl: %s - Err: %v", shortUrl, err))
	}

	return res
}

func GetUserUrls(userId string) []Url {
	var result []Url
	keys, err := storeService.redisClient.SMembers(ctx, "user:"+userId+":urls").Result()

	if err != nil {
		panic(fmt.Sprintf("Failed retrieving user URLs | user id: %s - Err: %v", userId, err))
	}

	for _, shortUrl := range keys {
		fullUrl, err := storeService.redisClient.Get(ctx, userId+":"+shortUrl).Result()

		if err != nil {
			panic(fmt.Sprintf("Failed retrieving full URL for shortUrl: %s | userId: %s - Err: %v", shortUrl, userId, err))
		}

		result = append(result, Url{
			ShortUrl: shortUrl,
			FullUrl:  fullUrl,
		})
	}

	return result
}

func DeleteShortUrl(shortUrl string, userId string) {
	keys := []string{shortUrl}
	keySet := "user:" + userId + ":urls"
	if len(strings.TrimSpace(shortUrl)) == 0 {
		urls := GetUserUrls(userId)
		for _, url := range urls {
			keys = append(keys, url.ShortUrl)
		}
	}

	if len(keys) > 0 {
		storeService.redisClient.SRem(ctx, keySet, keys)
	}

	for _, key := range keys {
		storeService.redisClient.Del(ctx, userId+":"+key)
	}

	remaining, _ := storeService.redisClient.SCard(ctx, keySet).Result()

	if remaining == 0 {
		storeService.redisClient.Del(ctx, keySet)
	}
}
