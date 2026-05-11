package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(URL string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Coudn't get DB instance")
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	if err != nil {
		log.Fatal("Coudn't run postgres")
	}
	return db, nil
}
