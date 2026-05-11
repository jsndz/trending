package handler

import (
	"strconv"

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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}
	articles, err := h.ArticleService.GetArticles(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, articles)
}
