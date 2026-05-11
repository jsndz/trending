wrk -t4 -c100 -d30s http://localhost:8080/api/v1/articles

Running 30s test @ http://localhost:8080/api/v1/articles
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   148.82ms  197.99ms 959.63ms   81.48%
    Req/Sec   411.44    325.42     1.90k    65.27%
  48982 requests in 30.09s, 232.51MB read
  Non-2xx or 3xx responses: 2669
Requests/sec:   1628.07
Transfer/sec:      7.73MB

There was many:
2026/05/11 14:58:15 /home/jaison/code/projects/trending/server/internal/repository/articles_repo.go:49 failed to connect to `user=developer database=trends`:
        [::1]:5432 (localhost): failed to receive message: unexpected EOF
        127.0.0.1:5432 (localhost): failed to receive message: unexpected EOF
[503.411ms] [rows:0] SELECT * FROM "articles" LIMIT 10
[GIN] 2026/05/11 - 14:58:15 | 500 | 503.47ms |             ::1 | GET      "/api/v1/articles"

2026/05/11 14:58:15 /home/jaison/code/projects/trending/server/internal/repository/articles_repo.go:49 SLOW SQL >= 200ms
[18230.672ms] [rows:0] SELECT * FROM "post_categories" WHERE "post_categories"."article_id" IN ('01KR1R65Z6TS16Q7GS4B7Z3HDR','01KR1R65Z6TTX2487TWN3NG0H4','01KR1R65Z6J4A2F5M374KMPTHK','01KR1R65Z6SES3NVXN96KB097B','01KR1R65Z6VKE9GC29S94B6P9P','01KR1R65Z6098SXPPXMNSS8A5S','01KR1R65Z64PSGTYM9H5H409V6','01KR1R65Z6WD0YD3W39X2PFPSV','01KR1R65Z6EWY34ZGC2QEMS21W','01KR1R65Z65EK5WAKYJP8ZBZY8')


These errors are caused by:

1. Unconfigured connection pool 
2. DB opened repeatedly


Since i don't see any DB opening repeatedly Problem is with connection pool configuration.
After the connection pool fix.

wrk -t4 -c100 -d30s http://localhost:8080/api/v1/articles
Running 30s test @ http://localhost:8080/api/v1/articles
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.54ms    3.94ms  66.64ms   72.93%
    Req/Sec     2.39k   128.96     2.68k    80.25%
  286040 requests in 30.02s, 1.40GB read
Requests/sec:   9527.49
Transfer/sec:     47.65MB