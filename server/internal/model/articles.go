package model

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Source string

const (
	SourceTechCrunch Source = "techcrunch"
	SourceHackerNews Source = "hackernews"
	SourceReuters    Source = "reuters"
	SourceBBC        Source = "bbc"
)

type Article struct {
	ID          string     `gorm:"primaryKey;size:26" json:"id"`
	Title       string     `gorm:"notNull" json:"title"`
	PublishedAt time.Time  `gorm:"notNull" json:"publishedAt"`
	Link        string     `gorm:"notNull;uniqueIndex" json:"link"`
	Author      string     `json:"author"`
	Category    []Category `gorm:"many2many:post_categories;constraint:OnDelete:CASCADE;" json:"category"`
	// when table is joined post_categories will be created and FK is automatically created on auto migrate
	Description string `json:"description"`
	Source      Source `json:"source"`
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	entropy := ulid.Monotonic(rand.Reader, 0)
	a.ID = ulid.MustNew(
		ulid.Timestamp(time.Now()),
		entropy,
	).String()

	return nil
}
