package dtos

import "time"

type BaseModel struct {
	ID        uint      `gorm:"primaryKey"`     // Primary key for all models
	CreatedAt time.Time `gorm:"autoCreateTime"` // Automatically set the created time
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // Automatically update the time on model update
}
