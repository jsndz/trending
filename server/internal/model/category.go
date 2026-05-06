package model

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Category struct {
	ID   string `gorm:"primaryKey;size:26"`
	Name string `gorm:"notNull"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	entropy := ulid.Monotonic(rand.Reader, 0)
	c.ID = ulid.MustNew(
		ulid.Timestamp(time.Now()),
		entropy,
	).String()

	return nil
}
