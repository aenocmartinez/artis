package usecase

import (
	"artis/src/museologo/domain/campos"
	"log"
)

type QuitarSubcampoACampoUseCase struct {
	repository campos.CampoRepository
}

func (q *QuitarSubcampoACampoUseCase) SetRepository(repository campos.CampoRepository) {
	q.repository = repository
}

func (q *QuitarSubcampoACampoUseCase) Ejecutar(idCampo, idSubcampo int64) bool {

	campo := campos.BuscarCampoPorId(q.repository, idCampo)
	if !campo.Existe() {
		log.Println("el campo no existe")
		return false
	}

	if !campo.EsCompuesto() {
		log.Println("el campo no es compuesto")
		return false
	}

	subcampo := campos.BuscarCampoPorId(q.repository, idSubcampo)
	if !subcampo.Existe() {
		log.Println("el subcampo no existe")
		return false
	}

	campoCompuesto, _ := campo.(*campos.CampoCompuesto)
	campoCompuesto.SetRepository(q.repository)
	return campoCompuesto.QuitarSubcampo(subcampo)
}
