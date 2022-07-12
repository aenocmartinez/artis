package usecase

import (
	"artis/src/museologo/domain/campos"
	"log"
)

type VerCampoUseCase struct {
	repository campos.CampoRepository
}

func (vc *VerCampoUseCase) SetRepository(repository campos.CampoRepository) {
	vc.repository = repository
}

func (vc *VerCampoUseCase) Ejecutar(idCampo int64) campos.DtoCampo {
	var dtoCampo campos.DtoCampo
	campo := campos.BuscarCampoPorId(vc.repository, idCampo)
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
