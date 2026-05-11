package bootstrap

import (
	"github.com/jsndz/trending/internal/feed/techcrunch"
	"github.com/jsndz/trending/internal/handler"
	"github.com/jsndz/trending/internal/repository"
	"github.com/jsndz/trending/internal/service"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Worker struct {
	FeedService *service.FeedService
}
type API struct {
	ArticleHandler *handler.ArticleHandler
}

func InitWorker(db *gorm.DB) *Worker {
	articlesRepo := repository.NewArticlesRepository(db)
	techcrunchProvider := techcrunch.NewTechCrunch()
	feedService := service.NewFeedSeervice(articlesRepo, techcrunchProvider)
	return &Worker{
		FeedService: feedService,
	}
}

func InitAPI(db *gorm.DB, redis *redis.Client) *API {
	articlesRepo := repository.NewArticlesRepository(db)
	articleService := service.NewArticleService(articlesRepo)
	articleHandler := handler.NewArticleHandler(articleService)
	return &API{
		ArticleHandler: articleHandler,
	}
}
