package formrequest

type GuardarColeccionFormRequest struct {
	Id     int64  `json:"id"`
	Nombre string `json:"nombre" binding:"required"`
}
