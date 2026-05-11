package service

import (
	"context"

	"github.com/jsndz/trending/internal/feed"
	"github.com/jsndz/trending/internal/repository"
	redisPkg "github.com/jsndz/trending/pkg/redis"
	"github.com/redis/go-redis/v9"
)

type FeedService struct {
	ArticlesRepo *repository.ArticlesRepository
	redis        *redis.Client
	Provider     feed.Provider
}

func NewFeedSeervice(articleRepo *repository.ArticlesRepository, provider feed.Provider, redis *redis.Client) *FeedService {
	return &FeedService{
		ArticlesRepo: articleRepo,
		Provider:     provider,
		redis:        redis,
	}
}

func (s *FeedService) SyncFeed(ctx context.Context) error {
	articles, err := s.Provider.Parser()
	if err != nil {
		return err
	}
	redisPkg.InvalidateTrendingCache(ctx, s.redis)
	err = s.ArticlesRepo.BatchCreate(articles)
	return err
}
