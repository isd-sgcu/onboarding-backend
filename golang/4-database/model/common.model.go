package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;size:191"` // ID field in JSON is id, and it is a primary key in the database
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp;autoCreateTime:nano"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp;autoUpdateTime:nano"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;type:timestamp"`
}

// https://gorm.io/docs/hooks.html
func (m *Base) BeforeCreate(_ *gorm.DB) error {
	m.ID = uuid.New()

	return nil
}
