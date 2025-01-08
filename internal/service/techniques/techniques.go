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

type TechniquesResponseService struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Points int    `json:"points"`
}

//Definição da interface para techniques service

type ITechniquesService interface {
	Create(techniques CreateTechniquesRequest)
	ListAllTechniques() []TechniquesResponseService
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

func (t *TechniquesServiceImpl) Create(techniques CreateTechniquesRequest) {
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

func (t *TechniquesServiceImpl) ListAllTechniques() []TechniquesResponseService {
	//receber tecnicas do repositorio
	techniquesFromRepo := t.TechniquesRepository.ListAllTechniques()

	//declarar variavel para ser preenchida com os dados vindos do repositorio
	var techniques []TechniquesResponseService

	//Fazer a iteração para cada tecnica que for retornada dentro da struct technique model.
	for _, value := range techniquesFromRepo {
		technique := TechniquesResponseService{
			ID:     value.ID,
			Name:   value.Name,
			Type:   value.Type,
			Points: value.Points,
		}

		techniques = append(techniques, technique)
	}
	//retorna a variavel do tipo ResponseService para a camada de controller pode acessar.
	return techniques
}
