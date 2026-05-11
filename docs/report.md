# Performance Optimization Improvements

## Initial State

Initial load testing showed severe database instability under concurrent traffic.

Benchmark:

```text id="s72qqk"
wrk -t4 -c100 -d30s http://localhost:8080/api/v1/articles
```

Results:

```text 
Requests/sec: 1628.07
Latency: ~148ms average
500 responses observed
```

Observed Issues:

* PostgreSQL connection exhaustion
* Slow queries
* Unexpected EOF database errors
* High latency under concurrent load

Example errors:

```text
failed to receive message: unexpected EOF
```

and:

```text
SLOW SQL >= 200ms
```

---

# 1. Database Connection Pool Optimization

## Problem

Database connections were not properly pooled, causing:

* connection exhaustion
* expensive connection creation
* unstable behavior under load

## Solution

Configured PostgreSQL connection pool limits:

* max open connections
* max idle connections
* connection lifetime management

## Result

Benchmark after pool optimization:

```text id="6g7rww"
Requests/sec: 9527.49
Latency: ~10ms average
```

Improvement:

* ~6x throughput increase
* stable database behavior
* elimination of connection-related failures

---

# 2. Redis Cache-Aside Caching

## Problem

Frequently requested article pages repeatedly queried PostgreSQL even when data rarely changed.

## Solution

Implemented Redis cache-aside strategy:

* check Redis first
* fallback to DB on cache miss
* populate Redis after DB read

Cache Key Example:

```text id="0r94hm"
articles:offset:0:limit:10
```

## Result

Benchmark after Redis caching:

```text id="ozm1bz"
Requests/sec: 20410.78
Latency: ~5ms average
```

Improvement:

* ~2x throughput increase over pooled DB setup
* major reduction in PostgreSQL read pressure
* significantly lower response latency

---

# 3. Request Coalescing / Cache Stampede Prevention

## Problem

When cache expired:

* many concurrent requests hit DB simultaneously
* duplicate expensive queries executed
* risk of DB overload during traffic spikes

## Solution

Implemented distributed request coalescing using Redis locks (`SETNX`).

Flow:

1. first request acquires lock
2. first request rebuilds cache
3. other requests wait for cache population
4. waiting requests reuse cached result

Lock Key Example:

```text
lock:articles:offset:0:limit:10
```

Additional protections:

* lock expiry timeout
* retry wait loop
* TTL jitter
* cache invalidation strategy

## Result

Benchmark after cache coalescing:

```text
Requests/sec: 17265.36
Latency: ~5.8ms average
```

Although raw throughput decreased slightly, this tradeoff is expected because:

* locking introduces coordination overhead
* retries add minimal latency
* synchronization increases CPU/network work

However, the primary objective was resilience, not maximum benchmark throughput.

Benefits:

* prevents cache stampede
* prevents duplicate DB work
* protects PostgreSQL during cache expiry events
* improves stability during burst traffic

---

# Additional Optimizations

## TTL Jitter

Added randomized cache expiry to avoid synchronized cache invalidation.

Example:

```text id="0rk75q"
60s + random(0-30s)
```

Purpose:

* prevents cache avalanche
* spreads DB load over time

---

## Negative Caching

Implemented short-lived caching for nonexistent records.

Purpose:

* prevents repeated DB lookups for invalid IDs
* reduces abuse/spam traffic impact

---

# Final Architecture

Current system includes:

* PostgreSQL connection pooling
* Redis cache-aside caching
* Distributed request coalescing
* TTL jitter
* Negative caching
* Cache invalidation on ingestion

---

# Final Performance Summary

| Stage              | Requests/sec | Avg Latency |
| ------------------ | ------------ | ----------- |
| Initial            | ~1.6k        | ~148ms      |
| Connection Pooling | ~9.5k        | ~10ms       |
| Redis Cache        | ~20.4k       | ~5ms        |
| Cache Coalescing   | ~17.2k       | ~5.8ms      |

---
