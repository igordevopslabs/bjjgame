package fighterscontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	fightersservice "github.com/igordevopslabs/bjjgame/internal/service/fighters"
	"github.com/igordevopslabs/bjjgame/pkg/helper"
)

type FightersResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

//Definie qa struct com campo para acessar a interface de service

type FightersController struct {
	fighterService fightersservice.IFightersService
}

// Inicia o construtor
func NewFightersController(svc fightersservice.IFightersService) *FightersController {
	return &FightersController{
		fighterService: svc,
	}
}

func (c *FightersController) FindAll(ctx *gin.Context) {
	fighterResponse := c.fighterService.FindAll()
	webResponse := FightersResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   fighterResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *FightersController) Create(ctx *gin.Context) {
	createFighterReq := fightersservice.CreateFightersRequest{}
	err := ctx.ShouldBindJSON(&createFighterReq)
	helper.ErrorPanic(err)

	c.fighterService.Create(createFighterReq)

	response := FightersResponse{
		Code:   200,
		Status: "Ok",
		Data:   createFighterReq,
	}

	ctx.JSON(http.StatusOK, response)

}
