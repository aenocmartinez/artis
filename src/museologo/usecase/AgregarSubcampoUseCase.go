package usecase

import (
	"artis/src/museologo/domain"
	"errors"
	"log"
)

type AgregarSubcampoUseCase struct {
	repository domain.MuseologoRepository
}

func (asc *AgregarSubcampoUseCase) SetRepository(repository domain.MuseologoRepository) {
	asc.repository = repository
}

func (asc *AgregarSubcampoUseCase) Ejecutar(idCampo, idSubcampo, orden int64) error {
	var err error

	campo := domain.BuscarCampoPorId(asc.repository, idCampo)
	if !campo.Existe() {
		err = errors.New("el campo no existe")
		log.Println("el campo no existe")
		return err
	}

	subcampo := domain.BuscarCampoPorId(asc.repository, idSubcampo)
	if !subcampo.Existe() {
		log.Println("el subcampo no existe")
		err = errors.New("el subcampo no existe")
		return err
	}

	campoCompuesto := domain.InstanceCampoCompuesto()
	campoCompuesto.SetId(idCampo)

	campoSimple := domain.InstanceCampoSimple()
	campoSimple.SetId(idSubcampo)

	campoCompuesto.SetRepository(asc.repository)
	campoCompuesto.AgregarSubcampo(*campoSimple, orden)

	return nil
}
