package usecase

import (
	"artis/src/museologo/domain/campos"
	"log"
)

type ListarSubcamposUseCase struct {
	repository campos.CampoRepository
}

func (ls *ListarSubcamposUseCase) SetRepository(repository campos.CampoRepository) {
	ls.repository = repository
}

func (ls *ListarSubcamposUseCase) Ejecutar(idCampo int64) []campos.DtoCampo {

	campo := campos.BuscarCampoPorId(ls.repository, idCampo)
	if !campo.Existe() {
		log.Println("el campo no existe")
		return []campos.DtoCampo{}
	}

	if !campo.EsCompuesto() {
		log.Println("el campo no es compuesto")
		return []campos.DtoCampo{}
	}

	campoCompuesto := campos.InstanceCampoCompuesto()
	campoCompuesto.SetId(idCampo)
	campoCompuesto.SetRepository(ls.repository)
	return campoCompuesto.ListarSubcampos()
}
