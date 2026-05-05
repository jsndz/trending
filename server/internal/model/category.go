package model

type Category struct {
	ID   string `gorm:"primaryKey;size:26"`
	Name string `gorm:"notNull"`
}
