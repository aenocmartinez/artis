package coleccion

type ColeccionRepository interface {
	Crear(coleccion Coleccion) bool
	BuscarColeccionPorId(id int64) Coleccion
	ListarColecciones() []DtoColeccion
	EliminarColeccion(id int64) bool
	ActualizarColeccion(coleccion Coleccion) bool
	AgregarCampo(idColeccion int64, campo CampoColeccion) bool
}
