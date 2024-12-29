package models

import "gorm.io/gorm"

type Fighters struct {
	ID uint `gorm:"primaryKey"`
	gorm.Model
	Name       string       `json:"name"`
	Team       string       `json:"team"`
	Style      string       `json:"style"`
	Techniques []Techniques `json:"techniques" gorm:"foreignKey:FighterID"`
	Overall    int64        `json:"overall"`
}
