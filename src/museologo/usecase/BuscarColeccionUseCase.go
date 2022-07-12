package usecase

import "artis/src/museologo/domain/coleccion"

type BuscarColeccionUseCase struct {
	repository coleccion.ColeccionRepository
}

func (bc *BuscarColeccionUseCase) SetRepository(repository coleccion.ColeccionRepository) {
	bc.repository = repository
}

func (bc *BuscarColeccionUseCase) Ejecutar(id int64) coleccion.Coleccion {
	return coleccion.BuscarColeccionPorId(bc.repository, id)
}
