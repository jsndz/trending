Cache Aside:
If exist in cache get 
else get from DB and update

```go

func (s *ArticleService) GetArticles(ctx context.Context, page int, limit int) ([]model.Article, error) {
	offset := (page - 1) * limit
	key := "articles:offset:" + strconv.Itoa(offset) + ":limit:" + strconv.Itoa(limit)
	val, err := s.redis.Get(ctx, key).Result()
	var articlesCache StoreData

	if err != nil && err == redis.Nil {
		// handle
		articles, err := s.ArticlesRepo.GetPaginated(limit, (page-1)*limit)
		if err != nil {
			return nil, err
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
		if err := s.redis.Set(ctx, key, articlesJSON, time.Second*60).Err(); err != nil {
			return nil, err
		}
		return articles, nil
	}
	json.Unmarshal([]byte(val), &articlesCache)

	return articlesCache.Data, nil
}


```