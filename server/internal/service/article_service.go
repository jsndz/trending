package service

import (
	"github.com/jsndz/trending/internal/model"
	"github.com/jsndz/trending/internal/repository"
)

type ArticleService struct {
	ArticlesRepo *repository.ArticlesRepository
}

func NewArticleService(articleRepo *repository.ArticlesRepository) *ArticleService {
	return &ArticleService{
		ArticlesRepo: articleRepo,
	}
}

func (s *ArticleService) GetArticles() ([]model.Article, error) {
	articles, err := s.ArticlesRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return articles, nil
}
