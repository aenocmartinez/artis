package domain

type InterfaceCampo interface {
	Id() int64
	Crear() bool
	Existe() bool
	Nombre() string
	Eliminar() bool
	Actualizar() bool
	Abreviatura() string
	Descripcion() string
	EsCompuesto() bool
	SetRepository(repository MuseologoRepository)
	SetId(id int64)
	SetNombre(nombre string)
	SetDescripcion(descripcion string)
	SetAbreviatura(abreviatura string)
}

type Campo struct {
	repository  MuseologoRepository
	id          int64  `json:"id"`
	nombre      string `json:"nombre"`
	descripcion string `json:"descripcion"`
	abreviatura string `json:"abreviatura"`
}

func (c *Campo) SetRepository(repository MuseologoRepository) {
	c.repository = repository
}

func (c *Campo) SetId(id int64) {
	c.id = id
}

func (c *Campo) SetNombre(nombre string) {
	c.nombre = nombre
}

func (c *Campo) SetDescripcion(descripcion string) {
	c.descripcion = descripcion
}

func (c *Campo) SetAbreviatura(abreviatura string) {
	c.abreviatura = abreviatura
}

func (c *Campo) Id() int64 {
	return c.id
}

func (c *Campo) Nombre() string {
	return c.nombre
}

func (c *Campo) Descripcion() string {
	return c.descripcion
}

func (c *Campo) Abreviatura() string {
	return c.abreviatura
}

func ListarCampos(repository MuseologoRepository) []DtoCampo {
	return repository.ListarCampos()
}

func BuscarCampoPorNombre(repository MuseologoRepository, nombre string) InterfaceCampo {
	return repository.BuscarCampoPorNombre(nombre)
}

func BuscarCampoPorId(repository MuseologoRepository, idCampo int64) InterfaceCampo {
	return repository.BuscarCampoPorId(idCampo)
}
