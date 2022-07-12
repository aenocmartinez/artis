package usecase

import "artis/src/museologo/domain"

type ListarCamposUseCase struct {
	repository domain.MuseologoRepository
}

func (lc *ListarCamposUseCase) SetRepository(repository domain.MuseologoRepository) {
	lc.repository = repository
}

func (lc *ListarCamposUseCase) Ejecutar() []domain.DtoCampo {
	return domain.ListarCampos(lc.repository)
}
