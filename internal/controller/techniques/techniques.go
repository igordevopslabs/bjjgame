package techniquescontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	techniquesservice "github.com/igordevopslabs/bjjgame/internal/service/techniques"
	"github.com/igordevopslabs/bjjgame/pkg/helper"
)

// Criar a struct de TechniquesResponse
type TechniquesResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// Criar a struct para acessar a Interface de techniques Service
type TechniquesController struct {
	techniqueService techniquesservice.ITechniquesService
}

//Criar o construtor de controller

func NewTechniquesController(svc techniquesservice.ITechniquesService) *TechniquesController {
	return &TechniquesController{
		techniqueService: svc,
	}
}

// Criar a função Creare Tecniques do Controller
func (t TechniquesController) Create(ctx *gin.Context) {
	createTechniquesReq := techniquesservice.CreateTechniquesRequest{}
	err := ctx.ShouldBindJSON(&createTechniquesReq)
	helper.ErrorPanic(err)

	t.techniqueService.Create(createTechniquesReq)

	response := TechniquesResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, response)
}
