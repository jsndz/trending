package bootstrap

import (
	"github.com/jsndz/trending/internal/feed/techcrunch"
	"github.com/jsndz/trending/internal/repository"
	"github.com/jsndz/trending/internal/service"
	"gorm.io/gorm"
)

type Worker struct {
	FeedService *service.FeedService
}

func InitWorker(db *gorm.DB) *Worker {
	articlesRepo := repository.NewArticlesRepository(db)
	techcrunchProvider := techcrunch.NewTechCrunch()
	feedService := service.NewFeedSeervice(articlesRepo, techcrunchProvider)
	return &Worker{
		FeedService: feedService,
	}
}
