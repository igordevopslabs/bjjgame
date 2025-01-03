package techniquesrepo

import (
	techniquesmodel "github.com/igordevopslabs/bjjgame/internal/model/techniques"
	"github.com/igordevopslabs/bjjgame/pkg/helper"
	"gorm.io/gorm"
)

//Definição da interface com os metodos

type ITechniquesRepo interface {
	Create(techniques techniquesmodel.Techniques)
}

type TechniquesRepoImpl struct {
	Db *gorm.DB
}

// instancia o construtor
func NewTechniquesRepoImpl(Db *gorm.DB) ITechniquesRepo {
	return &TechniquesRepoImpl{Db: Db}
}

func (t TechniquesRepoImpl) Create(techniques techniquesmodel.Techniques) {
	result := t.Db.Create(&techniques)
	helper.ErrorPanic(result.Error)
}

//paramos no service
