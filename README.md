**Project: Distributed “Trending Content” Aggregator**

**Objective**
Continuously ingest publicly available content, compute a time-decayed “trending” score using optimized database queries, and serve low-latency APIs backed by distributed caching. System must handle high read load, frequent updates, and cache consistency trade-offs.

---

**Data Sources (periodic fetch)**

* RSS feeds from BBC, Reuters, TechCrunch
* Optional engagement signals via Reddit API or GitHub API
  Polling interval: 1–5 minutes

---

**High-level Architecture**

* Ingestion workers (cron/scheduler) → normalize and store content
* Primary DB (Postgres) for durable storage and optimized queries
* Cache layer using Redis (clustered)
* Stateless API servers (horizontally scalable)
* Optional message bus for invalidation events

---

**Data Model**

* `posts(id, title, source, url, created_at)`
* `events(post_id, type, weight, created_at)` (views/clicks/upvotes)
* Optional precomputed table/materialized view for “trending”

Indexes:

* `(post_id, created_at)` on events
* `(created_at)` for time filtering
* Consider time-based partitioning on `events`

---

**Trending Computation (DB layer)**

* Score = weighted sum of events with exponential time decay
* Window: last 24 hours (configurable)
* Return top N (e.g., 50)

Optimizations:

* Materialized view refreshed every few minutes
* Partial indexes on recent data
* Read replicas for scaling reads
* Avoid full-table scans via time-bounded queries

---

**Caching Strategy**

* Pattern: cache-aside
* Keys: `trending:{region}:{category}`
* TTL: 60–120 seconds with jitter
* L1 (in-process, very short TTL ~5s) + L2 (Redis)

Read flow:

1. Check L1 → L2 → DB/materialized view
2. On miss, compute/fetch and populate caches

---

**Advanced Cache Controls**

* **Request coalescing:** single recomputation per key using Redis lock
* **Staggered TTL:** add random jitter to avoid synchronized expiry
* **Event-driven invalidation:** after ingestion or score update, delete affected keys
* **Negative caching:** cache empty results briefly
* **Hot key protection:** isolate “global trending” with tighter controls and L1 cache

---

**Consistency Model**

* Eventual consistency (acceptable for trending)
* Bounded staleness via TTL + periodic refresh
* Invalidation on significant updates to reduce staleness

---

**APIs**

* `GET /trending?region=&category=` → returns top items
* `GET /posts/{id}` → details
* `POST /event` (optional) → record interactions

---

**Scaling Strategy**

* API: horizontal scaling behind load balancer
* Redis: cluster/sharding for high QPS
* DB: primary + read replicas; partition `events` by time
* Separate ingestion workers from API path

---

**Failure Handling**

* Redis unavailable: fallback to DB (rate-limited)
* DB slow: serve stale cache if present
* Ingestion lag: rely on last computed cache
* Lock timeout: fail open (serve stale)

---

**Metrics & Observability**

* Cache hit rate (L1/L2)
* P50/P95/P99 latency
* Recompute frequency and duration
* DB query time and rows scanned
* Error rates (lock contention, timeouts)

---

**Security & Limits**

* Respect API/RSS rate limits
* Input sanitization for ingested content
* Basic rate limiting on public endpoints (can also use Redis)

---

