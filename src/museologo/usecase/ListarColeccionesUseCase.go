package usecase

import "artis/src/museologo/domain/coleccion"

type ListarColeccionesUseCase struct {
	repository coleccion.ColeccionRepository
}

func (lc *ListarColeccionesUseCase) SetRepository(repository coleccion.ColeccionRepository) {
	lc.repository = repository
}

func (lc *ListarColeccionesUseCase) Ejecutar() []coleccion.DtoColeccion {
	return coleccion.ListarColecciones(lc.repository)
}
