package usecase

import "artis/src/museologo/domain/campos"

type BuscarCampoUseCase struct {
	repository campos.CampoRepository
}

func (b *BuscarCampoUseCase) SetRepository(repository campos.CampoRepository) {
	b.repository = repository
}

func (b *BuscarCampoUseCase) Ejecutar(nombre string) campos.InterfaceCampo {
	return campos.BuscarCampoPorNombre(b.repository, nombre)
}
