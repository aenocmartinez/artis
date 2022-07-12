package usecase

import (
	"artis/src/museologo/domain"
	"log"
)

type QuitarSubcampoACampoUseCase struct {
	repository domain.MuseologoRepository
}

func (q *QuitarSubcampoACampoUseCase) SetRepository(repository domain.MuseologoRepository) {
	q.repository = repository
}

func (q *QuitarSubcampoACampoUseCase) Ejecutar(idCampo, idSubcampo int64) bool {

	campo := domain.BuscarCampoPorId(q.repository, idCampo)
	if !campo.Existe() {
		log.Println("el campo no existe")
		return false
	}

	if !campo.EsCompuesto() {
		log.Println("el campo no es compuesto")
		return false
	}

	subcampo := domain.BuscarCampoPorId(q.repository, idSubcampo)
	if !subcampo.Existe() {
		log.Println("el subcampo no existe")
		return false
	}

	campoCompuesto, _ := campo.(*domain.CampoCompuesto)
	campoCompuesto.SetRepository(q.repository)
	return campoCompuesto.QuitarSubcampo(subcampo)
}
