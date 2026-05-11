package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/jsndz/trending/internal/model"
	"github.com/jsndz/trending/internal/repository"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const NotFound = "NOT_FOUND"

type ArticleService struct {
	ArticlesRepo *repository.ArticlesRepository
	redis        *redis.Client
}

type StoreData struct {
	Data       []model.Article `json:"data"`
	NextOffset int             `json:"next_offset"`
	HasMore    bool            `json:"has_more"`
}

func NewArticleService(articleRepo *repository.ArticlesRepository, redis *redis.Client) *ArticleService {
	return &ArticleService{
		ArticlesRepo: articleRepo,
		redis:        redis,
	}
}

func (s *ArticleService) GetArticles(ctx context.Context, page int, limit int) ([]model.Article, error) {
	offset := (page - 1) * limit
	key := "articles:offset:" + strconv.Itoa(offset) + ":limit:" + strconv.Itoa(limit)
	val, err := s.redis.Get(ctx, key).Result()
	if val == NotFound {
		return nil, gorm.ErrRecordNotFound
	}
	var articlesCache StoreData

	if err == redis.Nil {
		// handle
		lock := "lock:" + key

		ok, err := s.redis.SetNX(ctx, lock, "1", time.Second*10).Result()
		// set if not exist of ok that means there is not db asking for this
		if err != nil {
			return nil, err
		}

		if ok {
			defer s.redis.Del(ctx, lock)
			log.Println("cache miss", key)
			articles, err := s.ArticlesRepo.GetPaginated(limit, (page-1)*limit)
			if err != nil {
				return nil, err
			}
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err := s.redis.Set(
					ctx,
					key,
					NotFound,
					5*time.Second,
				).Err()

				if err != nil {
					return nil, err
				}

				return nil, gorm.ErrRecordNotFound
			}
			hasMore := len(articles) > limit

			if hasMore {
				articles = articles[:limit]
			}
			articlesJSON, err := json.Marshal(&StoreData{
				Data:       articles,
				NextOffset: offset + limit,
				HasMore:    hasMore,
			})
			if err != nil {
				return nil, err
			}
			baseTTL := 60 * time.Second
			jitter := time.Duration(rand.Intn(30)) * time.Second

			ttl := baseTTL + jitter
			if err := s.redis.Set(ctx, key, articlesJSON, ttl).Err(); err != nil {
				return nil, err
			}
			return articles, nil
		}

		base := 50 * time.Millisecond

		for i := 0; i < 5; i++ {
			val, err := s.redis.Get(ctx, key).Result()
			if err == nil {
				var cache StoreData
				json.Unmarshal([]byte(val), &articlesCache)
				log.Println("cache hit", key)
				return cache.Data, nil
			}
			backoff := base * time.Duration(1<<i)
			jitter := time.Duration(rand.Intn(100)) * time.Millisecond
			time.Sleep(backoff + jitter)
		}
	}
	json.Unmarshal([]byte(val), &articlesCache)
	log.Println("cache hit", key)
	return articlesCache.Data, nil
}
