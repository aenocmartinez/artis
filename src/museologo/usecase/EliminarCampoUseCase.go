package usecase

import (
	"artis/src/museologo/domain/campos"
	"log"
)

type EliminarCampoUseCase struct {
	repository campos.CampoRepository
}

func (e *EliminarCampoUseCase) SetRepository(repository campos.CampoRepository) {
	e.repository = repository
}

func (e *EliminarCampoUseCase) Ejecutar(idCampo int64) bool {
	campo := campos.BuscarCampoPorId(e.repository, idCampo)
	if !campo.Existe() {
		log.Println("el campo no existe")
		return false
	}
	campo.SetRepository(e.repository)
	return campo.Eliminar()
}
