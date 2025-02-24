package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertAndQuery(t *testing.T) {
	initialLink := "https://www.google.com"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "2dDEQAS1"

	SaveUrlMap(shortUrl, initialLink, userUUId)
	retrievedUrl := GetFullUrl(shortUrl, userUUId)

	assert.Equal(t, initialLink, retrievedUrl)
}