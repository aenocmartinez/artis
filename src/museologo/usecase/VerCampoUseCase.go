package usecase

import (
	"artis/src/museologo/domain"
	"log"
)

type VerCampoUseCase struct {
	repository domain.CampoRepository
}

func (vc *VerCampoUseCase) SetRepository(repository domain.CampoRepository) {
	vc.repository = repository
}

func (vc *VerCampoUseCase) Ejecutar(idCampo int64) domain.DtoCampo {
	var dtoCampo domain.DtoCampo
	campo := domain.BuscarCampoPorId(vc.repository, idCampo)
	if !campo.Existe() {
		log.Println("el campo no existe")
		return dtoCampo
	}

	dtoCampo.Id = campo.Id()
	dtoCampo.Nombre = campo.Nombre()
	dtoCampo.Abreviatura = campo.Abreviatura()
	dtoCampo.Descripcion = campo.Descripcion()
	dtoCampo.EsCompuesto = campo.EsCompuesto()

	return dtoCampo
}
