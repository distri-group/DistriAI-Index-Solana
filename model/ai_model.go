package model

import "time"

type AiModel struct {
	Id        uint   `gorm:"primarykey"`
	Owner     string `gorm:"size:44;not null;index:idx_ai_models_owner_name"`
	Name      string `gorm:"size:50;not null;index:idx_ai_models_owner_name"`
	Framework uint8  `gorm:"not null"`
	License   uint8  `gorm:"not null"`
	Type1     uint32 `gorm:"not null"`
	Type2     uint32 `gorm:"not null"`
	Tags      string `gorm:"size:128;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
