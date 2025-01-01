package techniquesmodel

type Techniques struct {
	ID     int    `gorm:"type:int,primary_key"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Points int    `json:"points"`
}
