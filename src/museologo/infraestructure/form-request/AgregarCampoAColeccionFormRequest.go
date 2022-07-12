package formrequest

type AgregarCampoAColeccionFormRequest struct {
	EsObligatorio bool  `json:"esObligatorio"`
	EsUnico       bool  `json:"esUnico"`
	EsEditable    bool  `json:"esEditable"`
	Orden         int   `json:"orden"`
	IdColeccion   int64 `json:"idColeccion"`
	IdCampo       int64 `json:"idCampo"`
	IdLista       int64 `json:"idLista"`
	IdSecuencia   int64 `json:"idSecuencia"`
}
