package coleccion

type Coleccion struct {
	repository ColeccionRepository
	Id         int64
	Nombre     string
}

func (c *Coleccion) SetRepository(repository ColeccionRepository) {
	c.repository = repository
}

func (c *Coleccion) Crear() bool {
	return c.repository.Crear(*c)
}

func (c *Coleccion) Eliminar() bool {
	return c.repository.EliminarColeccion(c.Id)
}

func (c *Coleccion) Actualizar() bool {
	return c.repository.ActualizarColeccion(*c)
}

func (c *Coleccion) Existe() bool {
	return c.Id > 0
}

func (c *Coleccion) AgregarCampo(dtoCampoColeccion DtoCampoColeccion) bool {
	campoColeccion := InstanceCampoColeccion(dtoCampoColeccion.Campo, dtoCampoColeccion.Orden)
	campoColeccion.EsEditable = dtoCampoColeccion.EsEditable
	campoColeccion.EsUnico = dtoCampoColeccion.EsUnico
	campoColeccion.EsObligatorio = dtoCampoColeccion.EsObligatorio

	return c.repository.AgregarCampo(c.Id, *campoColeccion)
}

func BuscarColeccionPorId(repository ColeccionRepository, id int64) Coleccion {
	return repository.BuscarColeccionPorId(id)
}

func ListarColecciones(repository ColeccionRepository) []DtoColeccion {
	return repository.ListarColecciones()
}
