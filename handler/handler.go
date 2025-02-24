package handler

import (
	"net/http"

	"github.com/AndreaPallotta/url_shortener/shortener"
	"github.com/AndreaPallotta/url_shortener/store"
	"github.com/gin-gonic/gin"
)

type CreateUrlReq struct {
	FullUrl string `json:"full_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var req CreateUrlReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(req.FullUrl, req.UserId)
	store.SaveUrlMap(shortUrl, req.FullUrl, req.UserId)

	host := "http://localhost:8080/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("short_url")
	userId := c.DefaultQuery("user_id", "")
	fullUrl := store.GetFullUrl(shortUrl, userId)
	c.Redirect(302, fullUrl)
}

func GetUserUrls(c *gin.Context) {
	userId := c.Param("user_id")
	urls := store.GetUserUrls(userId)
	c.JSON(200, gin.H{
		"user_id": userId,
		"urls":    urls,
	})
}

func DeleteShortUrl(c *gin.Context) {
	userId := c.Param("user_id")
	shortUrl := c.Param("short_url")

	store.DeleteShortUrl(shortUrl, userId)
	c.JSON(200, gin.H{
		"message": "Short URL(s) deleted successfully",
	})
}