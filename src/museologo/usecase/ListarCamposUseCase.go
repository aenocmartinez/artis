package usecase

import "artis/src/museologo/domain"

type ListarCamposUseCase struct {
	repository domain.CampoRepository
}

func (lc *ListarCamposUseCase) SetRepository(repository domain.CampoRepository) {
	lc.repository = repository
}

func (lc *ListarCamposUseCase) Ejecutar() []domain.DtoCampo {
	return domain.ListarCampos(lc.repository)
}
