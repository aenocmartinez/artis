package coleccion

import "artis/src/museologo/domain/campos"

type CampoColeccion struct {
	campo          campos.InterfaceCampo
	Orden          int
	EsEditable     bool
	EsUnico        bool
	EsObligatorio  bool
	TieneSecuencia bool
	TieneLista     bool
}

func InstanceCampoColeccion(campo campos.InterfaceCampo, orden int) *CampoColeccion {
	return &CampoColeccion{
		campo: campo,
		Orden: orden,
	}
}
