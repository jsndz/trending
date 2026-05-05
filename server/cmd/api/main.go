package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jsndz/trending/internal/config"
	"github.com/jsndz/trending/pkg/db"
	"github.com/jsndz/trending/pkg/redis"
)

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	cfg := config.Load()
	database, err := db.InitDB(cfg.DBConnectURL)
	db.MigrateDB(database)
	if err != nil {
		panic(err)
	}
	_ = redis.NewRedisClient()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
