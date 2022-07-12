package usecase

import (
	"artis/src/museologo/domain/coleccion"
	formrequest "artis/src/museologo/infraestructure/form-request"
	"log"
)

type ActualizarColeccionUseCase struct {
	repository coleccion.ColeccionRepository
}

func (ac *ActualizarColeccionUseCase) SetRepository(repository coleccion.ColeccionRepository) {
	ac.repository = repository
}

func (ac *ActualizarColeccionUseCase) Ejecutar(req formrequest.GuardarColeccionFormRequest) bool {

	coleccion := coleccion.BuscarColeccionPorId(ac.repository, req.Id)
	if !coleccion.Existe() {
		log.Println("la coleccion no existe")
		return false
	}

	coleccion.SetRepository(ac.repository)
	coleccion.Nombre = req.Nombre
	coleccion.Id = req.Id

	return coleccion.Actualizar()

}
