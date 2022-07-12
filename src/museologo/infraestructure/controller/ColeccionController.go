package controller

import (
	"artis/src/museologo/domain/campos"
	"artis/src/museologo/domain/coleccion"
	formrequest "artis/src/museologo/infraestructure/form-request"
	"artis/src/museologo/usecase"
	"errors"

	"github.com/gin-gonic/gin"
)

type ColeccionController struct {
	coleccionRepository coleccion.ColeccionRepository
	campoRepository     campos.CampoRepository
}

func InstanceColeccionController(coleccionRepository coleccion.ColeccionRepository, campoRepository campos.CampoRepository) *ColeccionController {
	return &ColeccionController{
		coleccionRepository: coleccionRepository,
		campoRepository:     campoRepository,
	}
}

func (ctr *ColeccionController) CrearColeccion(c *gin.Context) error {
	var req formrequest.GuardarColeccionFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return err
	}

	crearColeccionUseCase := usecase.CrearColeccionUseCase{}
	crearColeccionUseCase.SetColeccionRepository(ctr.coleccionRepository)
	crearColeccionUseCase.SetCampoRepository(ctr.campoRepository)
	respuesta := crearColeccionUseCase.Ejecutar(req)
	if !respuesta {
		return errors.New("ha ocurrido un error")
	}

	return nil
}
