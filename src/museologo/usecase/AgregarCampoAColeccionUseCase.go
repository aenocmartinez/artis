package usecase

import (
	"artis/src/museologo/domain/campos"
	"artis/src/museologo/domain/coleccion"
	formrequest "artis/src/museologo/infraestructure/form-request"
	"artis/src/museologo/infraestructure/repository/mysql"
	"log"
)

type AgregarCampoAColeccionUseCase struct {
	repository coleccion.ColeccionRepository
}

func (ac *AgregarCampoAColeccionUseCase) SetRepository(repository coleccion.ColeccionRepository) {
	ac.repository = repository
}

func (ac *AgregarCampoAColeccionUseCase) Ejecutar(req formrequest.AgregarCampoAColeccionFormRequest) bool {

	oColeccion := coleccion.BuscarColeccionPorId(ac.repository, req.IdColeccion)
	if !oColeccion.Existe() {
		log.Println("la coleccion no existe")
		return false
	}

	oCampo := campos.BuscarCampoPorId(mysql.InstanceCampoDB(), req.IdCampo)
	if !oCampo.Existe() {
		log.Println("el campo no existe")
		return false
	}

	dtoCampoColeccion := coleccion.DtoCampoColeccion{
		IdLista:       req.IdLista,
		IdSecuencia:   req.IdSecuencia,
		Orden:         req.Orden,
		EsObligatorio: req.EsObligatorio,
		EsUnico:       req.EsUnico,
		EsEditable:    req.EsEditable,
		Campo:         oCampo,
	}

	return oColeccion.AgregarCampo(dtoCampoColeccion)
}
