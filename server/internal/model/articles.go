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
	ID          string    `gorm:"primaryKey;size:26"`
	Title       string    `gorm:"notNull"`
	PublishedAt time.Time `gorm:"notNull"`
	Link        string    `gorm:"notNull,unique"`
	Author      string
	Category    []Category `gorm:"many2many:post_categories;constraint:OnDelete:CASCADE;"`
	// when table is joined post_categories will be created and FK is automatically created on auto migrate
	Description string
	Source      Source
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	entropy := ulid.Monotonic(rand.Reader, 0)
	a.ID = ulid.MustNew(
		ulid.Timestamp(time.Now()),
		entropy,
	).String()

	return nil
}
