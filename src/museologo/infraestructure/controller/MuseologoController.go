package controller

import (
	"artis/src/museologo/domain/campos"
	formrequest "artis/src/museologo/infraestructure/form-request"
	"artis/src/museologo/usecase"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MuseologoController struct {
	repository campos.CampoRepository
}

func InstanceMuseologoController(repository campos.CampoRepository) *MuseologoController {
	return &MuseologoController{
		repository: repository,
	}
}

func (ctr *MuseologoController) CrearCampo(c *gin.Context) error {
	var req formrequest.GuardarCampoFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return err
	}

	crearCampoUseCase := usecase.CrearCampoUseCase{}
	crearCampoUseCase.SetRepository(ctr.repository)
	respuesta := crearCampoUseCase.Ejecutar(req)
	if !respuesta {
		return errors.New("ha ocurrido un error")
	}

	return nil
}

func (ctr *MuseologoController) ListarCampos(c *gin.Context) []campos.DtoCampo {
	listarCamposUseCase := usecase.ListarCamposUseCase{}
	listarCamposUseCase.SetRepository(ctr.repository)
	return listarCamposUseCase.Ejecutar()
}

func (ctr *MuseologoController) VerCampo(c *gin.Context) campos.DtoCampo {
	var err error
	var id int = 0

	if c.Param("id") != "" {
		id, err = strconv.Atoi(c.Param("id"))
		if err != nil {
			id = 0
		}
	}

	verCampoUseCase := usecase.VerCampoUseCase{}
	verCampoUseCase.SetRepository(ctr.repository)
	return verCampoUseCase.Ejecutar(int64(id))
}

func (ctr *MuseologoController) AgregarSubcampo(c *gin.Context) error {
	var req formrequest.AgregarSubcampoFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return err
	}

	agregarSubcampoUseCase := usecase.AgregarSubcampoUseCase{}
	agregarSubcampoUseCase.SetRepository(ctr.repository)
	return agregarSubcampoUseCase.Ejecutar(req.IdCampo, req.IdSubcampo, req.Orden)
}

func (ctr *MuseologoController) ListarSubcampos(c *gin.Context) []campos.DtoCampo {
	var err error
	var id int = 0

	if c.Param("id") != "" {
		id, err = strconv.Atoi(c.Param("id"))
		if err != nil {
			id = 0
		}
	}
	listarSubcamposUseCase := usecase.ListarSubcamposUseCase{}
	listarSubcamposUseCase.SetRepository(ctr.repository)
	return listarSubcamposUseCase.Ejecutar(int64(id))
}

func (ctr *MuseologoController) EliminarCampo(c *gin.Context) error {
	var err error
	var id int = 0

	if c.Param("id") != "" {
		id, err = strconv.Atoi(c.Param("id"))
		if err != nil {
			id = 0
		}
	}

	eliminarCampoUseCase := usecase.EliminarCampoUseCase{}
	eliminarCampoUseCase.SetRepository(ctr.repository)
	respuesta := eliminarCampoUseCase.Ejecutar(int64(id))
	if !respuesta {
		return errors.New("ha ocurrido un error")
	}

	return nil
}

func (ctr *MuseologoController) ActualizarCampo(c *gin.Context) error {
	var req formrequest.GuardarCampoFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return err
	}

	actualizarCampoUseCase := usecase.ActualizarCampoUseCase{}
	actualizarCampoUseCase.SetRepository(ctr.repository)
	respuesta := actualizarCampoUseCase.Ejecutar(req)
	if !respuesta {
		return errors.New("ha ocurrido un error")
	}

	return nil
}

func (ctr *MuseologoController) QuitarSubcampo(c *gin.Context) error {
	var err error
	var idCampo int = 0
	var idSubcampo int = 0

	if c.Param("idCampo") != "" {
		idCampo, err = strconv.Atoi(c.Param("idCampo"))
		if err != nil {
			idCampo = 0
		}
	}

	if c.Param("idSubcampo") != "" {
		idSubcampo, err = strconv.Atoi(c.Param("idSubcampo"))
		if err != nil {
			idSubcampo = 0
		}
	}

	quitarSubcampoUseCase := usecase.QuitarSubcampoACampoUseCase{}
	quitarSubcampoUseCase.SetRepository(ctr.repository)
	respuesta := quitarSubcampoUseCase.Ejecutar(int64(idCampo), int64(idSubcampo))
	if !respuesta {
		return errors.New("ha ocurrido un error")
	}

	return nil
}
