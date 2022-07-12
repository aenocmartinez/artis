package usecase

import (
	"artis/src/museologo/domain/campos"
	"errors"
	"log"
)

type AgregarSubcampoUseCase struct {
	repository campos.CampoRepository
}

func (asc *AgregarSubcampoUseCase) SetRepository(repository campos.CampoRepository) {
	asc.repository = repository
}

func (asc *AgregarSubcampoUseCase) Ejecutar(idCampo, idSubcampo, orden int64) error {
	var err error

	campo := campos.BuscarCampoPorId(asc.repository, idCampo)
	if !campo.Existe() {
		err = errors.New("el campo no existe")
		log.Println("el campo no existe")
		return err
	}

	subcampo := campos.BuscarCampoPorId(asc.repository, idSubcampo)
	if !subcampo.Existe() {
		log.Println("el subcampo no existe")
		err = errors.New("el subcampo no existe")
		return err
	}

	campoCompuesto := campos.InstanceCampoCompuesto()
	campoCompuesto.SetId(idCampo)

	campoSimple := campos.InstanceCampoSimple()
	campoSimple.SetId(idSubcampo)

	campoCompuesto.SetRepository(asc.repository)
	campoCompuesto.AgregarSubcampo(*campoSimple, orden)

	return nil
}
