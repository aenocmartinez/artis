package formrequest

type GuardarCampoFormRequest struct {
	Id          int64  `json:"id"`
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion"`
	Abreviatura string `json:"abreviatura"`
	EsCompuesto bool   `json:"esCompuesto"`
}
