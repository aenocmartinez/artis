package usecase

import (
	"artis/src/museologo/domain/coleccion"
	"log"
)

type EliminarColeccionUseCase struct {
	repository coleccion.ColeccionRepository
}

func (ec *EliminarColeccionUseCase) SetRepository(repository coleccion.ColeccionRepository) {
	ec.repository = repository
}

func (ec *EliminarColeccionUseCase) Ejecutar(id int64) bool {

	coleccion := coleccion.BuscarColeccionPorId(ec.repository, id)
	if !coleccion.Existe() {
		log.Println("la coleccion no existe")
		return false
	}

	return coleccion.Eliminar()
}
