package formrequest

type AgregarSubcampoFormRequest struct {
	IdCampo    int64 `json:"idCampo" binding:"required"`
	IdSubcampo int64 `json:"idSubcampo" binding:"required"`
	Orden      int64 `json:"orden" binding:"required"`
}
