package usecase

import (
	"artis/src/museologo/domain/campos"
	formrequest "artis/src/museologo/infraestructure/form-request"
)

type CrearCampoUseCase struct {
	repository campos.CampoRepository
}

func (useCase *CrearCampoUseCase) SetRepository(repository campos.CampoRepository) {
	useCase.repository = repository
}

func (useCase *CrearCampoUseCase) Ejecutar(req formrequest.GuardarCampoFormRequest) bool {
	dtoCampo := campos.DtoCampo{
		Nombre:      req.Nombre,
		Descripcion: req.Descripcion,
		Abreviatura: req.Abreviatura,
		EsCompuesto: req.EsCompuesto,
	}
	campo := campos.CampoFactoria(dtoCampo)
	campo.SetRepository(useCase.repository)

	respuesta := campo.Crear()

	if campo.EsCompuesto() {
		agregarSubcampoUseCase := AgregarSubcampoUseCase{}
		agregarSubcampoUseCase.SetRepository(useCase.repository)
		agregarSubcampoUseCase.Ejecutar(campo.Id(), campo.Id(), 1)
	}

	return respuesta
}
