package usecase

import (
	"artis/src/museologo/domain/coleccion"
	"log"
)

type BuscarColeccionUseCase struct {
	repository coleccion.ColeccionRepository
}

func (bc *BuscarColeccionUseCase) SetRepository(repository coleccion.ColeccionRepository) {
	bc.repository = repository
}

func (bc *BuscarColeccionUseCase) Ejecutar(id int64) coleccion.DtoColeccion {
	oColeccion := coleccion.BuscarColeccionPorId(bc.repository, id)
	if !oColeccion.Existe() {
		log.Println("la coleccion no existe")
		return coleccion.DtoColeccion{}
	}

	return coleccion.DtoColeccion{
		Id:     oColeccion.Id,
		Nombre: oColeccion.Nombre,
	}
}
