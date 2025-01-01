package fightersmodel

type Fighters struct {
	ID      int    `gorm:"type:int,primary_key"`
	Name    string `json:"name"`
	Team    string `json:"team"`
	Style   string `json:"style"`
	Overall int    `json:"overall"`
}
