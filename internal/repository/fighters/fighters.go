package fightersrepo

import (
	"fmt"

	fightersmodel "github.com/igordevopslabs/bjjgame/internal/model/fighters"
	techniquesmodel "github.com/igordevopslabs/bjjgame/internal/model/techniques"
	"github.com/igordevopslabs/bjjgame/pkg/helper"
	"gorm.io/gorm"
)

//Definição da interface para ser acessada pelo data handler

type IFightersRepo interface {
	Create(fighters fightersmodel.Fighters)
	FindAll() []fightersmodel.Fighters
	FindFIghtersById(ids []int) ([]fightersmodel.Fighters, error)
	FindFIghtersBySingleId(id int) (fightersmodel.Fighters, error)
	UpdateFighter(fighters fightersmodel.Fighters)
	UpdateMatches(fighter fightersmodel.Fighters)
	FindTechniquesByIds(ids []int) ([]techniquesmodel.Techniques, error)
	FindTechniquesByBelt(belt string) ([]techniquesmodel.Techniques, error)
	UpdateFighterWithTechniques(fighter fightersmodel.Fighters) error
}

//Definição dos metodos para interagir com a camada de repository

type UpdateFightersRepo struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Team  string `json:"team"`
	Style string `json:"style"`
}

type UpdateFighterMatchesRepo struct {
	ID      int    `json:"id"`
	Matches int    `json:"matches"`
	Belt    string `json:"belt"`
	Overall int    `json:"overall"`
}

type FightersRepoImpl struct {
	Db *gorm.DB
}

func NewFighterRepoImpl(Db *gorm.DB) IFightersRepo {
	return &FightersRepoImpl{Db: Db}
}

// Lista todos os lutadores
func (f *FightersRepoImpl) FindAll() []fightersmodel.Fighters {
	var fighters []fightersmodel.Fighters
	// Carregar as técnicas associadas com Preload
	result := f.Db.Preload("Techniques").Find(&fighters)
	helper.ErrorPanic(result.Error)
	return fighters
}

// busca um par de lutadores no banco
func (f *FightersRepoImpl) FindFIghtersById(ids []int) ([]fightersmodel.Fighters, error) {
	var fighters []fightersmodel.Fighters
	result := f.Db.Preload("Techniques").Where("id in ?", ids).Find(&fighters)
	if result.Error != nil {
		return nil, result.Error
	}
	return fighters, result.Error
}

// busca por um unico ID
func (f *FightersRepoImpl) FindFIghtersBySingleId(id int) (fightersmodel.Fighters, error) {
	var fighter fightersmodel.Fighters
	// Carregar as técnicas associadas com Preload
	result := f.Db.Preload("Techniques").First(&fighter, id)
	if result.Error != nil {
		return fightersmodel.Fighters{}, result.Error
	}
	return fighter, nil
}

// Cria os lutadores no Repository
func (f *FightersRepoImpl) Create(fighters fightersmodel.Fighters) {
	result := f.Db.Create(&fighters)
	helper.ErrorPanic(result.Error)
}

func (f *FightersRepoImpl) UpdateFighter(fighters fightersmodel.Fighters) {
	updateFighter := UpdateFightersRepo{
		ID:    fighters.ID,
		Name:  fighters.Name,
		Style: fighters.Style,
		Team:  fighters.Team,
	}

	result := f.Db.Model(&fighters).Updates(updateFighter)
	helper.ErrorPanic(result.Error)
}

func (f *FightersRepoImpl) UpdateMatches(fighter fightersmodel.Fighters) {
	updateFighterMatch := UpdateFighterMatchesRepo{
		ID:      fighter.ID,
		Matches: fighter.Matches,
		Belt:    fighter.Belt,
		Overall: fighter.Overall,
	}

	result := f.Db.Model(&fighter).Updates(updateFighterMatch)
	helper.ErrorPanic(result.Error)
}

func (f *FightersRepoImpl) UpdateTechniquesForFighter(fighterId int, techniques []techniquesmodel.Techniques) error {
	result := f.Db.Model(&fightersmodel.Fighters{ID: fighterId}).Association("Techniques").Replace(techniques)
	return result
}

func (f *FightersRepoImpl) FindTechniquesByIds(ids []int) ([]techniquesmodel.Techniques, error) {
	var techniques []techniquesmodel.Techniques
	result := f.Db.Where("id IN ?", ids).Find(&techniques)
	if result.Error != nil {
		return nil, result.Error
	}
	return techniques, nil
}

func (f *FightersRepoImpl) FindTechniquesByBelt(belt string) ([]techniquesmodel.Techniques, error) {
	var techniques []techniquesmodel.Techniques
	result := f.Db.Where("required_belt = ?", belt).Find(&techniques)
	if result.Error != nil {
		return nil, result.Error
	}
	return techniques, nil
}

func (f *FightersRepoImpl) UpdateFighterWithTechniques(fighter fightersmodel.Fighters) error {
	// Associar as técnicas ao lutador
	for i := range fighter.Techniques {
		fighter.Techniques[i].FighterID = fighter.ID
	}

	// Log antes de atualizar as associações
	fmt.Printf("Techniques to associate: %+v\n", fighter.Techniques)

	// Atualizar as associações de técnicas
	err := f.Db.Model(&fighter).Association("Techniques").Replace(fighter.Techniques)
	if err != nil {
		return err
	}

	// Persistir as alterações do lutador no banco
	result := f.Db.Save(&fighter)
	if result.Error != nil {
		return result.Error
	}

	// Log após salvar o lutador
	fmt.Printf("Updated fighter: %+v\n", fighter)

	return nil
}
