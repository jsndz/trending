cache stampede is when all the keys suddenly expire and all request suddenly hit the DB
which can be very dangerous like if suddenly 1000 req hit DB

Request Coalescing solves this by allowing only ONE request fetches from DB
others wait

so this can be achieved using redis set NX
set when not exist 
if the key for lock not exist set
so that only one request can handle it 
set the lock 
make db query and cache is rebuild 

other req can use the same cache

dont forget to delete the lock  

```go 


func (s *ArticleService) GetArticles(ctx context.Context, page int, limit int) ([]model.Article, error) {
	offset := (page - 1) * limit
	key := "articles:offset:" + strconv.Itoa(offset) + ":limit:" + strconv.Itoa(limit)
	val, err := s.redis.Get(ctx, key).Result()
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


```

TTL jitter
when one request is doing the db request others have to wait
we are using retry with exponential backoff + jitter

Add jitter for expiry
Many different keys
expiring simultaneously

Negative caching:

if data does not exist on DB and malicous attacker request that data again and again 
so cache will call DB and too many can hit
so if the record is not found cache is as null
so that it can directly return nil
usually has low ttl like 5sec since the status of data might change
like does not exist -> exist


Invalidation:

invalidate cache when new data is uploaded