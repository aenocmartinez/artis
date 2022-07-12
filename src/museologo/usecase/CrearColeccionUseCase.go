package usecase

import (
	"artis/src/museologo/domain/campos"
	"artis/src/museologo/domain/coleccion"
	formrequest "artis/src/museologo/infraestructure/form-request"
)

type CrearColeccionUseCase struct {
	coleccionRepository coleccion.ColeccionRepository
	campoRepository     campos.CampoRepository
}

func (cc *CrearColeccionUseCase) SetColeccionRepository(repository coleccion.ColeccionRepository) {
	cc.coleccionRepository = repository
}

func (cc *CrearColeccionUseCase) SetCampoRepository(repository campos.CampoRepository) {
	cc.campoRepository = repository
}

func (cc *CrearColeccionUseCase) Ejecutar(req formrequest.GuardarColeccionFormRequest) bool {

	coleccion := coleccion.Coleccion{
		Nombre: req.Nombre,
	}

	coleccion.SetRepository(cc.coleccionRepository)

	return coleccion.Crear()
}
