package bootstrap

import (
	"github.com/jsndz/trending/internal/repository"
	"github.com/jsndz/trending/internal/service"
	"gorm.io/gorm"
)

type Worker struct {
	FeedService *service.FeedService
}

func InitWorker(db *gorm.DB) *Worker {
	articlesRepo := repository.NewArticlesRepository(db)
	feedService := service.NewFeedSeervice(articlesRepo)
	return &Worker{
		FeedService: feedService,
	}
}
