package main

import (
	"fmt"

	"github.com/AndreaPallotta/url_shortener/handler"
	"github.com/AndreaPallotta/url_shortener/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HEY",
		})
	})

	r.GET("/user/:user_id/urls", handler.GetUserUrls)
	r.GET("/:short_url", handler.HandleShortUrlRedirect)
	r.POST("/generate", handler.CreateShortUrl)
	r.DELETE("/user/:user_id/urls/:short_url", handler.DeleteShortUrl)
	r.DELETE("/user/:user_id/urls", handler.DeleteShortUrl)

	store.InitStore()

	err := r.Run(":8080")
	fmt.Println("Server started on port [8080]")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server - Error: %v", err))
	}
}
