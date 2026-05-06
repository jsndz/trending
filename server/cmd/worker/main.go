package main

import (
	"log"

	"github.com/jsndz/trending/internal/bootstrap"
	"github.com/jsndz/trending/internal/config"
	"github.com/jsndz/trending/internal/scheduler"
	"github.com/jsndz/trending/pkg/db"
)

func main() {
	cfg := config.Load()
	database, err := db.InitDB(cfg.DBConnectURL)
	db.MigrateDB(database)
	if err != nil {
		panic(err)
	}
	worker := bootstrap.InitWorker(database)

	err = scheduler.Schedule(func() {
		err = worker.FeedService.SyncFeed()
		if err != nil {
			log.Println(err)
		}
	})

	if err != nil {
		log.Fatal(err)
	}

	select {}

}
