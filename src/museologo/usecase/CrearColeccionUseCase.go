package usecase

import (
	"artis/src/museologo/domain/coleccion"
	formrequest "artis/src/museologo/infraestructure/form-request"
)

type CrearColeccionUseCase struct {
	repository coleccion.ColeccionRepository
}

func (cc *CrearColeccionUseCase) SetRepository(repository coleccion.ColeccionRepository) {
	cc.repository = repository
}

func (cc *CrearColeccionUseCase) Ejecutar(req formrequest.GuardarColeccionFormRequest) bool {

	coleccion := coleccion.Coleccion{
		Nombre: req.Nombre,
	}

	coleccion.SetRepository(cc.repository)

	return coleccion.Crear()
}
