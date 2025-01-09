package techniquesmodel

type Techniques struct {
	ID            int    `gorm:"type:int,primary_key;autoIncrement"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Points        int    `json:"points"`
	Required_Belt string `json:"required_belt"`
	FighterID     int    `gorm:"index" json:"fighter_id"` // Chave estrangeira

}
