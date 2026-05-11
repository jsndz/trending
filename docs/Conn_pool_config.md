Connection pool is number of connection provided for connecting with DB.
Creating a new connection for every request is expensive:

TCP setup
authentication
DB session creation

So create connection pool

every request
→ get connection from pool
→ execute query
→ return connection to pool

too many connections: postgres overwhelmed
too less connection: slow reply


```go

	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Coudn't get DB instance")
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

```
