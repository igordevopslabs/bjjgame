package fightersrepo

import (
	"errors"

	fightersmodel "github.com/igordevopslabs/bjjgame/internal/model/fighters"
	"github.com/igordevopslabs/bjjgame/pkg/helper"
	"gorm.io/gorm"
)

//Definição da interface para ser acessada pelo data handler

type IFightersRepo interface {
	Create(fighters fightersmodel.Fighters)
	FindAll() []fightersmodel.Fighters
	FindFIghtersById(ids []int) ([]fightersmodel.Fighters, error)
	FindFIghtersBySingleId(id int) ([]fightersmodel.Fighters, error)
}

//Definição dos metodos para interagir com a camada de repository

//Struct para acessar as opções do Gorm

type FightersRepoImpl struct {
	Db *gorm.DB
}

func NewFighterRepoImpl(Db *gorm.DB) IFightersRepo {
	return &FightersRepoImpl{Db: Db}
}

// Lista todos os lutadores
func (f FightersRepoImpl) FindAll() []fightersmodel.Fighters {
	var fighters []fightersmodel.Fighters
	result := f.Db.Find(&fighters)
	helper.ErrorPanic(result.Error)
	return fighters
}

// busca um par de lutadores no banco
func (f FightersRepoImpl) FindFIghtersById(ids []int) ([]fightersmodel.Fighters, error) {
	var fighters []fightersmodel.Fighters
	result := f.Db.Where("id in ?", ids).Find(&fighters)
	if result.Error != nil {
		return nil, result.Error
	}
	return fighters, result.Error
}

// busca por um unico ID
func (f FightersRepoImpl) FindFIghtersBySingleId(id int) ([]fightersmodel.Fighters, error) {
	var fighters []fightersmodel.Fighters
	result := f.Db.Find(&fighters, id).Find(&fighters)
	if result.Error != nil {
		return nil, result.Error
	} else if result != nil {
		return fighters, nil
	} else {
		return nil, errors.New("fighter is not found")
	}
}

// Cria os lutadores no Repository
func (f FightersRepoImpl) Create(fighters fightersmodel.Fighters) {
	result := f.Db.Create(&fighters)
	helper.ErrorPanic(result.Error)
}
