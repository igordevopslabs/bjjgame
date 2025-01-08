package techniquesrepo

import (
	techniquesmodel "github.com/igordevopslabs/bjjgame/internal/model/techniques"
	"github.com/igordevopslabs/bjjgame/pkg/helper"
	"gorm.io/gorm"
)

//Definição da interface com os metodos

type ITechniquesRepo interface {
	Create(techniques techniquesmodel.Techniques)
	ListAllTechniques() []techniquesmodel.Techniques
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

func (t *TechniquesRepoImpl) ListAllTechniques() []techniquesmodel.Techniques {
	var techniques []techniquesmodel.Techniques
	result := t.Db.Find(&techniques)
	helper.ErrorPanic(result.Error)
	return techniques
}
