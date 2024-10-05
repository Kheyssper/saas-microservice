package models

import (
	"time"

	"gorm.io/gorm"
)

type Platform struct {
	gorm.Model
	PlatformName string    `gorm:"type:varchar(255)" json:"platform_name"`
	PlatformSlug string    `gorm:"type:varchar(255)" json:"platform_slug"`
	CreatorID    int       `json:"creator_id"`
	Status       string    `gorm:"type:varchar(20)" json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
