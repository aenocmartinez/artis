package coleccion

import "artis/src/museologo/domain/campos"

type DtoCampoColeccion struct {
	IdLista       int64
	IdSecuencia   int64
	EsObligatorio bool
	EsUnico       bool
	EsEditable    bool
	Orden         int
	Campo         campos.InterfaceCampo
}
