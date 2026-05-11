package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jsndz/trending/internal/bootstrap"
	"github.com/jsndz/trending/internal/config"
	"github.com/jsndz/trending/pkg/db"
	"github.com/jsndz/trending/pkg/redis"
)

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	cfg := config.Load()
	database, err := db.InitDB(cfg.DBConnectURL)
	db.MigrateDB(database)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewRedisClient()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	app := bootstrap.InitAPI(database, redisClient)
	api := router.Group("/api/v1")
	api.GET("/articles", app.ArticleHandler.GetArticles)
	router.Run(":8080")
}
