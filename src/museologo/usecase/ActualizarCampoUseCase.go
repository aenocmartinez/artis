package usecase

import (
	"artis/src/museologo/domain"
	formrequest "artis/src/museologo/infraestructure/form-request"
	"log"
)

type ActualizarCampoUseCase struct {
	repository domain.CampoRepository
}

func (a *ActualizarCampoUseCase) SetRepository(repository domain.CampoRepository) {
	a.repository = repository
}

func (a *ActualizarCampoUseCase) Ejecutar(req formrequest.GuardarCampoFormRequest) bool {

	campo := domain.BuscarCampoPorId(a.repository, req.Id)
	if !campo.Existe() {
		log.Println("el campo no existe")
		return false
	}

	campo.SetRepository(a.repository)
	campo.SetNombre(req.Nombre)
	campo.SetAbreviatura(req.Abreviatura)
	campo.SetDescripcion(req.Descripcion)

	return campo.Actualizar()
}
