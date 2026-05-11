package repository

import (
	"github.com/jsndz/trending/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ArticlesRepository struct {
	db *gorm.DB
}

func NewArticlesRepository(db *gorm.DB) *ArticlesRepository {
	return &ArticlesRepository{
		db: db,
	}
}

func (r *ArticlesRepository) Create(article *model.Article) error {
	return r.db.Create(article).Error
}

func (r *ArticlesRepository) Get(id string) (*model.Article, error) {
	var article model.Article
	err := r.db.Preload("Category").First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticlesRepository) BatchCreate(articles *[]model.Article) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "link"}},
		DoNothing: true}).CreateInBatches(articles, len(*articles)).Error
}

func (r *ArticlesRepository) GetAll() ([]model.Article, error) {
	var articles []model.Article
	err := r.db.Preload("Category").Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *ArticlesRepository) GetPaginated(limit, offset int) ([]model.Article, error) {
	var articles []model.Article
	err := r.db.Preload("Category").Limit(limit + 1).Offset(offset).Find(&articles).Error
	return articles, err
}
