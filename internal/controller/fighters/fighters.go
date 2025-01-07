package fighterscontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	fightersrepo "github.com/igordevopslabs/bjjgame/internal/repository/fighters"
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

func (c *FightersController) FindById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	helper.ErrorPanic(err)

	result, err := c.fighterService.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	webResponse := FightersResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (c *FightersController) FightersCompare(ctx *gin.Context) {
	//receber os ids via parametro da request
	idParam1 := ctx.Param("id1")
	idParam2 := ctx.Param("id2")

	id1, err := strconv.Atoi(idParam1)
	helper.ErrorPanic(err)
	id2, err := strconv.Atoi(idParam2)
	helper.ErrorPanic(err)

	result, err := c.fighterService.FightersOverallCompare(id1, id2)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := FightersResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

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

func (c *FightersController) UpdateFighter(ctx *gin.Context) {
	updateFighterReq := fightersrepo.UpdateFightersRepo{}
	err := ctx.ShouldBindJSON(&updateFighterReq)

	helper.ErrorPanic(err)
	fighterId := ctx.Param("id")

	id, err := strconv.Atoi(fighterId)
	helper.ErrorPanic(err)

	updateFighterReq.ID = id

	c.fighterService.UpdateFighter(updateFighterReq)

	webResponse := FightersResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   "Data Uptaded",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
