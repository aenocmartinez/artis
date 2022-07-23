package controller

import (
	"artis/src/museologo/domain/campos"
	"artis/src/museologo/domain/coleccion"
	formrequest "artis/src/museologo/infraestructure/form-request"
	"artis/src/museologo/usecase"
	"errors"
	"strconv"

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

func (ctr *ColeccionController) ListarColecciones(c *gin.Context) []coleccion.DtoColeccion {
	listarColeccionesUseCase := usecase.ListarColeccionesUseCase{}
	listarColeccionesUseCase.SetRepository(ctr.coleccionRepository)
	return listarColeccionesUseCase.Ejecutar()
}

func (ctr *ColeccionController) BuscarColeccion(c *gin.Context) coleccion.DtoColeccion {
	var err error
	var id int = 0

	if c.Param("id") != "" {
		id, err = strconv.Atoi(c.Param("id"))
		if err != nil {
			id = 0
		}
	}

	buscarColeccionUseCase := usecase.BuscarColeccionUseCase{}
	buscarColeccionUseCase.SetRepository(ctr.coleccionRepository)
	return buscarColeccionUseCase.Ejecutar(int64(id))
}

func (ctr *ColeccionController) ActualizarColeccion(c *gin.Context) error {
	var req formrequest.GuardarColeccionFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return err
	}

	actualizarColeccionUseCase := usecase.ActualizarColeccionUseCase{}
	actualizarColeccionUseCase.SetRepository(ctr.coleccionRepository)
	respuesta := actualizarColeccionUseCase.Ejecutar(req)
	if !respuesta {
		return errors.New("ha ocurrido un error")
	}

	return nil
}
