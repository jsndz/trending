package service

import (
	"io"
	"net/http"

	"github.com/jsndz/trending/internal/feed/techcrunch"
	"github.com/jsndz/trending/internal/repository"
)

type FeedService struct {
	ArticlesRepo *repository.ArticlesRepository
}

func NewFeedSeervice(articleRepo *repository.ArticlesRepository) *FeedService {
	return &FeedService{
		ArticlesRepo: articleRepo,
	}
}

func (s *FeedService) SyncFeed() error {
	resp, err := http.Get("https://techcrunch.com/feed/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	articles, err := techcrunch.Parser(data)
	if err != nil {
		return err
	}
	err = s.ArticlesRepo.BatchCreate(articles)
	return err
}
