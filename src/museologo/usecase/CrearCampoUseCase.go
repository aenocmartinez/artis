package usecase

import (
	"artis/src/museologo/domain"
	formrequest "artis/src/museologo/infraestructure/form-request"
)

type CrearCampoUseCase struct {
	repository domain.CampoRepository
}

func (useCase *CrearCampoUseCase) SetRepository(repository domain.CampoRepository) {
	useCase.repository = repository
}

func (useCase *CrearCampoUseCase) Ejecutar(req formrequest.GuardarCampoFormRequest) bool {
	dtoCampo := domain.DtoCampo{
		Nombre:      req.Nombre,
		Descripcion: req.Descripcion,
		Abreviatura: req.Abreviatura,
		EsCompuesto: req.EsCompuesto,
	}
	campo := domain.CampoFactoria(dtoCampo)
	campo.SetRepository(useCase.repository)

	respuesta := campo.Crear()

	if campo.EsCompuesto() {
		agregarSubcampoUseCase := AgregarSubcampoUseCase{}
		agregarSubcampoUseCase.SetRepository(useCase.repository)
		agregarSubcampoUseCase.Ejecutar(campo.Id(), campo.Id(), 1)
	}

	return respuesta
}
