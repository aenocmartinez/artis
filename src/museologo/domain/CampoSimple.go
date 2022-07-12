package domain

type CampoSimple struct {
	repository CampoRepository
	campo      *Campo `json:"campo"`
}

func InstanceCampoSimple() *CampoSimple {
	return &CampoSimple{
		campo: &Campo{},
	}
}

func (cs *CampoSimple) SetRepository(repository CampoRepository) {
	cs.repository = repository
}

func (cs *CampoSimple) EsCompuesto() bool {
	return false
}

func (cs *CampoSimple) Crear() bool {
	return cs.repository.CrearCampo(cs)
}

func (cs *CampoSimple) Eliminar() bool {
	return cs.repository.EliminarCampo(cs.Id())
}

func (cs *CampoSimple) Actualizar() bool {
	return cs.repository.ActualizarCampo(cs.campo)
}

func (cs *CampoSimple) Nombre() string {
	return cs.campo.Nombre()
}

func (cs *CampoSimple) Abreviatura() string {
	return cs.campo.Abreviatura()
}

func (cs *CampoSimple) Id() int64 {
	return cs.campo.Id()
}

func (cs *CampoSimple) Descripcion() string {
	return cs.campo.Descripcion()
}

//
func (cs *CampoSimple) SetNombre(nombre string) {
	cs.campo.SetNombre(nombre)
}

func (cs *CampoSimple) SetAbreviatura(abreviatura string) {
	cs.campo.SetAbreviatura(abreviatura)
}

func (cs *CampoSimple) SetId(id int64) {
	cs.campo.SetId(id)
}

func (cs *CampoSimple) SetDescripcion(descripcion string) {
	cs.campo.SetDescripcion(descripcion)
}

//

func (cs *CampoSimple) Existe() bool {
	return cs.Nombre() != ""
}
