package domain

type MuseologoRepository interface {
	CrearCampo(campo InterfaceCampo) bool
	EliminarCampo(id int64) bool
	ActualizarCampo(campo *Campo) bool
	ListarCampos() []DtoCampo
	BuscarCampoPorNombre(nombre string) InterfaceCampo
	BuscarCampoPorId(id int64) InterfaceCampo
	AgregarSubcampo(idCampo, idSubcampo, orden int64) bool
	QuitarSubcampo(idCampo, idSubcampo int64) bool
	ListarSubcampos(idCampo int64) []DtoCampo
}
