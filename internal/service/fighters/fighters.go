package fightersservice

import (
	"github.com/go-playground/validator/v10"
	fightersmodel "github.com/igordevopslabs/bjjgame/internal/model/fighters"
	fightersrepo "github.com/igordevopslabs/bjjgame/internal/repository/fighters"
	"github.com/igordevopslabs/bjjgame/pkg/helper"
)

type CreateFightersRequest struct {
	Name    string `validate:"required,min=1,max=200" json:"name"`
	Team    string `validate:"required,min=1,max=200" json:"team"`
	Style   string `validate:"required,min=1,max=200" json:"style"`
	Overall int    `validate:"required" json:"overall"`
}

type FightersResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Team    string `json:"team"`
	Style   string `json:"style"`
	Overall int    `json:"overall"`
}

// Definição da interface da camanda de serviço
type IFightersService interface {
	Create(fighters CreateFightersRequest)
	FindAll() []FightersResponse
}

type FightersServiceImpl struct {
	FighterRepository fightersrepo.IFightersRepo
	Validate          *validator.Validate
}

//Cria o construtor

func NewFightersServiceImpl(fighterRepository fightersrepo.IFightersRepo, validate *validator.Validate) IFightersService {
	return &FightersServiceImpl{
		FighterRepository: fighterRepository,
		Validate:          validate,
	}
}

func (f FightersServiceImpl) FindAll() []FightersResponse {
	//instancia o banco
	fightersFromRepo := f.FighterRepository.FindAll()

	var fighters []FightersResponse

	for _, value := range fightersFromRepo {
		fighter := FightersResponse{
			ID:      value.ID,
			Name:    value.Name,
			Team:    value.Team,
			Style:   value.Style,
			Overall: value.Overall,
		}
		fighters = append(fighters, fighter)

	}

	return fighters
}

func (f FightersServiceImpl) Create(fighters CreateFightersRequest) {
	//valida a struc recebida
	err := f.Validate.Struct(fighters)
	helper.ErrorPanic(err)
	fighterModel := fightersmodel.Fighters{
		Name:    fighters.Name,
		Team:    fighters.Team,
		Style:   fighters.Style,
		Overall: fighters.Overall,
	}

	f.FighterRepository.Create(fighterModel)
}
