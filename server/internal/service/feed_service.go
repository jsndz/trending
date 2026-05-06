package service

import (
	"github.com/jsndz/trending/internal/feed"
	"github.com/jsndz/trending/internal/repository"
)

type FeedService struct {
	ArticlesRepo *repository.ArticlesRepository
	Provider     feed.Provider
}

func NewFeedSeervice(articleRepo *repository.ArticlesRepository, provider feed.Provider) *FeedService {
	return &FeedService{
		ArticlesRepo: articleRepo,
		Provider:     provider,
	}
}

func (s *FeedService) SyncFeed() error {
	articles, err := s.Provider.Parser()
	if err != nil {
		return err
	}
	err = s.ArticlesRepo.BatchCreate(articles)
	return err
}
