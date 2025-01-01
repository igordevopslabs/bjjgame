package techniquesservice

import (
	"github.com/go-playground/validator/v10"
	techniquesmodel "github.com/igordevopslabs/bjjgame/internal/model/techniques"
	techniquesrepo "github.com/igordevopslabs/bjjgame/internal/repository/techniques"
	"github.com/igordevopslabs/bjjgame/pkg/helper"
)

type CreateTechniquesRequest struct {
	Name   string `validate:"required,min=1,max=200" json:"name"`
	Type   string `validate:"required,min=1,max=200" json:"type"`
	Points int    `validate:"required" json:"points"`
}

//Definição da interface para techniques service

type ITechniquesService interface {
	Create(techniques CreateTechniquesRequest)
}

type TechniquesServiceImpl struct {
	TechniquesRepository techniquesrepo.ITechniquesRepo
	Validate             *validator.Validate
}

//Inicia o construtor

func NewTechniquesServiceImpl(techniques techniquesrepo.ITechniquesRepo, validate *validator.Validate) ITechniquesService {
	return &TechniquesServiceImpl{
		TechniquesRepository: techniques,
		Validate:             validate,
	}
}

func (t TechniquesServiceImpl) Create(techniques CreateTechniquesRequest) {
	err := t.Validate.Struct(techniques)
	helper.ErrorPanic(err)

	//preencher o model
	techniquesModel := techniquesmodel.Techniques{
		Name:   techniques.Name,
		Type:   techniques.Type,
		Points: techniques.Points,
	}

	t.TechniquesRepository.Create(techniquesModel)
}
