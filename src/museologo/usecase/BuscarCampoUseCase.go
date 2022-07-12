package usecase

import "artis/src/museologo/domain"

type BuscarCampoUseCase struct {
	repository domain.MuseologoRepository
}

func (b *BuscarCampoUseCase) SetRepository(repository domain.MuseologoRepository) {
	b.repository = repository
}

func (b *BuscarCampoUseCase) Ejecutar(nombre string) domain.InterfaceCampo {
	return domain.BuscarCampoPorNombre(b.repository, nombre)
}
