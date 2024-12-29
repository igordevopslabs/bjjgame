package models

import "gorm.io/gorm"

type Techniques struct {
	ID uint `gorm:"primaryKey"`
	gorm.Model
	Name      string `json:"name"`
	Type      string `json:"type"`
	Points    int64  `json:"points"`
	FighterID uint   `json:"fighter_id"`
}
