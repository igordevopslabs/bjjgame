package fightersrepo

import (
	fightersmodel "github.com/igordevopslabs/bjjgame/internal/model/fighters"
	"github.com/igordevopslabs/bjjgame/pkg/helper"
	"gorm.io/gorm"
)

//Definição da interface para ser acessada pelo data handler

type IFightersRepo interface {
	Create(fighters fightersmodel.Fighters)
}

//Definição dos metodos para interagir com a camada de repository

//Struct para acessar as opções do Gorm

type FightersRepoImpl struct {
	Db *gorm.DB
}

func NewFighterRepoImpl(Db *gorm.DB) IFightersRepo {
	return &FightersRepoImpl{Db: Db}
}

func (f FightersRepoImpl) Create(fighters fightersmodel.Fighters) {
	result := f.Db.Create(&fighters)
	helper.ErrorPanic(result.Error)
}
