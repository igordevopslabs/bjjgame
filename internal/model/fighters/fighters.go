package fightersmodel

import techniquesmodel "github.com/igordevopslabs/bjjgame/internal/model/techniques"

type Fighters struct {
	ID         int                          `gorm:"type:int,primary_key;autoIncrement"`
	Name       string                       `json:"name"`
	Team       string                       `json:"team"`
	Style      string                       `json:"style"`
	Overall    int                          `json:"overall"`
	Matches    int                          `json:"matches"`
	Belt       string                       `json:"belt"`
	Techniques []techniquesmodel.Techniques `gorm:"foreignKey:FighterID" json:"techniques"` // Relacionamento 1:N

}
