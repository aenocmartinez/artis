package domain

type DtoCampo struct {
	Id          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Abreviatura string `json:"abreviatura"`
	EsCompuesto bool   `json:"compuesto"`
}
