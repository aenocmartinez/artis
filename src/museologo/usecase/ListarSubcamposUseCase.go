package usecase

import (
	"artis/src/museologo/domain"
	"log"
)

type ListarSubcamposUseCase struct {
	repository domain.MuseologoRepository
}

func (ls *ListarSubcamposUseCase) SetRepository(repository domain.MuseologoRepository) {
	ls.repository = repository
}

func (ls *ListarSubcamposUseCase) Ejecutar(idCampo int64) []domain.DtoCampo {

	campo := domain.BuscarCampoPorId(ls.repository, idCampo)
	if !campo.Existe() {
		log.Println("el campo no existe")
		return []domain.DtoCampo{}
	}

	if !campo.EsCompuesto() {
		log.Println("el campo no es compuesto")
		return []domain.DtoCampo{}
	}

	campoCompuesto := domain.InstanceCampoCompuesto()
	campoCompuesto.SetId(idCampo)
	campoCompuesto.SetRepository(ls.repository)
	return campoCompuesto.ListarSubcampos()
}
