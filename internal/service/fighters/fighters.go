package fightersservice

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	fightersmodel "github.com/igordevopslabs/bjjgame/internal/model/fighters"
	techniquesmodel "github.com/igordevopslabs/bjjgame/internal/model/techniques"
	fightersrepo "github.com/igordevopslabs/bjjgame/internal/repository/fighters"
	"github.com/igordevopslabs/bjjgame/pkg/helper"
)

type CreateFightersRequest struct {
	Name    string `validate:"required,min=1,max=200" json:"name"`
	Team    string `validate:"required,min=1,max=200" json:"team"`
	Style   string `validate:"required,min=1,max=200" json:"style"`
	Overall int    `validate:"required" json:"overall"`
	Matches int    `validate:"required" json:"matches"`
	Belt    string `validate:"required" json:"belt"`
}

type FightersResponse struct {
	ID         int                          `json:"id"`
	Name       string                       `json:"name"`
	Team       string                       `json:"team"`
	Style      string                       `json:"style"`
	Overall    int                          `json:"overall"`
	Matches    int                          `json:"matches"`
	Belt       string                       `json:"belt"`
	Techniques []techniquesmodel.Techniques `json:"techniques"`
}

// Definição da interface da camanda de serviço
type IFightersService interface {
	Create(fighters CreateFightersRequest)
	FindAll() []FightersResponse
	FightersOverallCompare(id1, id2 int) (string, error)
	FindById(id int) (FightersResponse, error)
	UpdateFighter(fighter fightersrepo.UpdateFightersRepo)
	UpdateFighterMatches(fighter fightersrepo.UpdateFighterMatchesRepo)
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

func (f *FightersServiceImpl) FindAll() []FightersResponse {
	//instancia o banco
	fightersFromRepo := f.FighterRepository.FindAll()

	var fighters []FightersResponse

	for _, value := range fightersFromRepo {
		fighter := FightersResponse{
			ID:         value.ID,
			Name:       value.Name,
			Team:       value.Team,
			Style:      value.Style,
			Overall:    value.Overall,
			Matches:    value.Matches,
			Belt:       value.Belt,
			Techniques: value.Techniques,
		}
		fighters = append(fighters, fighter)

	}

	return fighters
}

func (f *FightersServiceImpl) FindById(id int) (FightersResponse, error) {
	//instancia o banco
	fightersFromRepo, err := f.FighterRepository.FindFIghtersBySingleId(id)
	if err != nil {
		return FightersResponse{}, err
	}

	fighterResponse := FightersResponse{
		ID:      fightersFromRepo.ID,
		Name:    fightersFromRepo.Name,
		Team:    fightersFromRepo.Team,
		Style:   fightersFromRepo.Style,
		Overall: fightersFromRepo.Overall,
		Matches: fightersFromRepo.Matches,
		Belt:    fightersFromRepo.Belt,
	}

	return fighterResponse, nil

}

func (f *FightersServiceImpl) FightersOverallCompare(id1, id2 int) (string, error) {
	//buscar os lutadores retornados atraves do repository
	fighters, err := f.FighterRepository.FindFIghtersById([]int{id1, id2})
	if err != nil {
		return "", err
	}

	//verifica se existe os dois ids
	if len(fighters) != 2 {
		return "", fmt.Errorf("to have a fight, needs two different fighters")
	}

	//fazer a comparação dos ids.
	fighter1 := fighters[0]
	fighter2 := fighters[1]

	//soma Overall com os pontos das tecnicas
	totalOverallFighter1 := fighter1.Overall + f.SumFighterTechniquesPoints(fighter1.Techniques)
	totalOverallFighter2 := fighter2.Overall + f.SumFighterTechniquesPoints(fighter2.Techniques)

	if totalOverallFighter1 > totalOverallFighter2 {
		return fmt.Sprintf("'%s' Faixa: '%s' Overall:'%d' Wins '%s' Faixa: '%s' Overall:'%d", fighter1.Name, fighter1.Belt, totalOverallFighter1, fighter2.Name, fighter2.Belt, totalOverallFighter2), nil
	} else if totalOverallFighter2 > totalOverallFighter1 {
		return fmt.Sprintf("'%s' Faixa: '%s' Overall:'%d' Wins '%s' Faixa: '%s' Overall:'%d", fighter2.Name, fighter2.Belt, totalOverallFighter2, fighter1.Name, fighter1.Belt, totalOverallFighter1), nil
	} else {
		return "Draw", nil
	}
}

func (f *FightersServiceImpl) Create(fighters CreateFightersRequest) {
	//valida a struc recebida
	err := f.Validate.Struct(fighters)
	helper.ErrorPanic(err)
	fighterModel := fightersmodel.Fighters{
		Name:    fighters.Name,
		Team:    fighters.Team,
		Style:   fighters.Style,
		Overall: fighters.Overall,
		Matches: fighters.Matches,
		Belt:    fighters.Belt,
	}

	f.FighterRepository.Create(fighterModel)
}

func (f *FightersServiceImpl) UpdateFighter(fighter fightersrepo.UpdateFightersRepo) {
	fighterData, err := f.FighterRepository.FindFIghtersBySingleId(fighter.ID)
	helper.ErrorPanic(err)

	if fighterData.Name != "" {
		fighterData.Name = fighter.Name
	}

	if fighterData.Team != "" {
		fighterData.Team = fighter.Team
	}

	if fighterData.Style != "" {
		fighterData.Style = fighter.Style
	}

	f.FighterRepository.UpdateFighter(fighterData)
}

func (f *FightersServiceImpl) UpdateFighterMatches(fighter fightersrepo.UpdateFighterMatchesRepo) {
	fmt.Printf("UpdateFighterMatches called with fighter: %+v\n", fighter)

	fighterData, err := f.FighterRepository.FindFIghtersBySingleId(fighter.ID)
	helper.ErrorPanic(err)
	fmt.Printf("Fetched fighter data: %+v\n", fighterData)

	fighterData.Matches = fighter.Matches

	switch {
	case fighterData.Matches < 11:
		fighterData.Overall = 5
	case fighterData.Matches < 51:
		fighterData.Overall = 20
	case fighterData.Matches < 99:
		fighterData.Overall = 50
	case fighterData.Matches < 499:
		fighterData.Overall = 100
	case fighterData.Matches < 999:
		fighterData.Overall = 250
	case fighterData.Matches < 4999:
		fighterData.Overall = 1000
	}

	switch {
	case fighterData.Overall < 50:
		fighterData.Belt = "Branca"
	case fighterData.Overall < 100:
		fighterData.Belt = "Azul"
	case fighterData.Overall < 150:
		fighterData.Belt = "Roxa"
	case fighterData.Overall < 200:
		fighterData.Belt = "Marrom"
	case fighterData.Overall > 200:
		fighterData.Belt = "Preta"
	case fighterData.Overall > 500:
		fighterData.Belt = "Coral"
	case fighterData.Overall >= 1000:
		fighterData.Belt = "Vermelha"
	}

	fmt.Printf("Fighter's new belt: %s\n", fighterData.Belt)

	//buscar a quantidade de tecnicas permitidas com base na faixa
	//limit := f.getTechniqueLimitForBelt(fighter.Belt)

	techniques, err := f.FighterRepository.FindTechniquesByBelt(fighterData.Belt)
	helper.ErrorPanic(err)

	fmt.Printf("Techniques being associated: %+v\n", techniques)

	fighterData.Techniques = techniques

	err = f.FighterRepository.UpdateFighterWithTechniques(fighterData)
	helper.ErrorPanic(err)

	fmt.Printf("Updated fighter with techniques: %+v\n", fighterData)

}

func (f *FightersServiceImpl) SumFighterTechniquesPoints(techniques []techniquesmodel.Techniques) int {
	totalPoints := 0

	for _, technique := range techniques {
		totalPoints += technique.Points
	}

	return totalPoints
}
