package usecase

import (
	"artis/src/museologo/domain"
	"log"
)

type EliminarCampoUseCase struct {
	repository domain.CampoRepository
}

func (e *EliminarCampoUseCase) SetRepository(repository domain.CampoRepository) {
	e.repository = repository
}

func (e *EliminarCampoUseCase) Ejecutar(idCampo int64) bool {
	campo := domain.BuscarCampoPorId(e.repository, idCampo)
	if !campo.Existe() {
		log.Println("el campo no existe")
		return false
	}
	campo.SetRepository(e.repository)
	return campo.Eliminar()
}
