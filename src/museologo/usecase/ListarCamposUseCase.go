package usecase

import "artis/src/museologo/domain/campos"

type ListarCamposUseCase struct {
	repository campos.CampoRepository
}

func (lc *ListarCamposUseCase) SetRepository(repository campos.CampoRepository) {
	lc.repository = repository
}

func (lc *ListarCamposUseCase) Ejecutar() []campos.DtoCampo {
	return campos.ListarCampos(lc.repository)
}
