package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jsndz/trending/internal/bootstrap"
	"github.com/jsndz/trending/internal/config"
	"github.com/jsndz/trending/internal/scheduler"
	"github.com/jsndz/trending/pkg/db"
	"github.com/jsndz/trending/pkg/redis"
)

func main() {
	cfg := config.Load()
	database, err := db.InitDB(cfg.DBConnectURL)
	db.MigrateDB(database)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewRedisClient()

	worker := bootstrap.InitWorker(database, redisClient)

	err = scheduler.Schedule(func() {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			30*time.Second,
		)
		defer cancel()
		err = worker.FeedService.SyncFeed(ctx)
		if err != nil {
			log.Println(err)
		}
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("WORKER RUNNING")
	select {}

}
