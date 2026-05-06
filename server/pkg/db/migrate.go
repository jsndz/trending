package db

import (
	"github.com/jsndz/trending/internal/model"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&model.Article{}, &model.Category{})
}
