package domain

type CampoCompuesto struct {
	repository MuseologoRepository
	campo      *Campo  `json:"campo"`
	Subcampos  []Campo `json:"subcampos"`
}

func InstanceCampoCompuesto() *CampoCompuesto {
	return &CampoCompuesto{
		campo: &Campo{},
	}
}

func (cc *CampoCompuesto) SetRepository(repository MuseologoRepository) {
	cc.repository = repository
}

func (cc *CampoCompuesto) EsCompuesto() bool {
	return true
}

func (cc *CampoCompuesto) AgregarSubcampo(subcampo CampoSimple, orden int64) bool {
	return cc.repository.AgregarSubcampo(cc.Id(), subcampo.Id(), orden)
}

func (cc *CampoCompuesto) QuitarSubcampo(campo InterfaceCampo) bool {
	return cc.repository.QuitarSubcampo(cc.Id(), campo.Id())
}

func (cc *CampoCompuesto) ListarSubcampos() []DtoCampo {
	return cc.repository.ListarSubcampos(cc.Id())
}

func (cc *CampoCompuesto) TotalSubcampos() int {
	return len(cc.Subcampos)
}

func (cc *CampoCompuesto) Crear() bool {
	return cc.repository.CrearCampo(cc)
}

func (cc *CampoCompuesto) Eliminar() bool {
	return cc.repository.EliminarCampo(cc.Id())
}

func (cc *CampoCompuesto) Actualizar() bool {
	return cc.repository.ActualizarCampo(cc.campo)
}

func (cc *CampoCompuesto) Nombre() string {
	return cc.campo.Nombre()
}

func (cc *CampoCompuesto) Abreviatura() string {
	return cc.campo.Abreviatura()
}

func (cc *CampoCompuesto) Id() int64 {
	return cc.campo.Id()
}

func (cc *CampoCompuesto) Descripcion() string {
	return cc.campo.Descripcion()
}

func (cc *CampoCompuesto) Existe() bool {
	return cc.Nombre() != ""
}

func (cc *CampoCompuesto) SetId(id int64) {
	cc.campo.id = id
}

func (cc *CampoCompuesto) SetNombre(nombre string) {
	cc.campo.nombre = nombre
}

func (cc *CampoCompuesto) SetDescripcion(descripcion string) {
	cc.campo.descripcion = descripcion
}

func (cc *CampoCompuesto) SetAbreviatura(abreviatura string) {
	cc.campo.abreviatura = abreviatura
}
