package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jsndz/trending/internal/service"
)

type ArticleHandler struct {
	ArticleService *service.ArticleService
}

func NewArticleHandler(articleService *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		ArticleService: articleService,
	}
}

func (h *ArticleHandler) GetArticles(c *gin.Context) {
	articles, err := h.ArticleService.GetArticles()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, articles)
}
