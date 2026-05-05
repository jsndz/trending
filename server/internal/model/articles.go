package model

import "time"

type Source string

const (
	SourceTechCrunch Source = "techcrunch"
	SourceHackerNews Source = "hackernews"
	SourceReuters    Source = "reuters"
	SourceBBC        Source = "bbc"
)

type Article struct {
	ID          string    `gorm:"primaryKey;size:26"`
	Title       string    `gorm:"notNull"`
	PublishedAt time.Time `gorm:"notNull"`
	Author      string
	Category    []Category `gorm:"many2many:post_categories;constraint:OnDelete:CASCADE;"`
	// when table is joined post_categories will be created and FK is automatically created on auto migrate
	Description string
	Source      Source
}
