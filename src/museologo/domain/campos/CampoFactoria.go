package campos

func CampoFactoria(dtoCampo DtoCampo) InterfaceCampo {
	var campo InterfaceCampo = &CampoSimple{
		campo: &Campo{
			nombre:      dtoCampo.Nombre,
			descripcion: dtoCampo.Descripcion,
			abreviatura: dtoCampo.Abreviatura,
			id:          dtoCampo.Id,
		},
	}

	if dtoCampo.EsCompuesto {
		campo = &CampoCompuesto{
			campo: &Campo{
				nombre:      dtoCampo.Nombre,
				descripcion: dtoCampo.Descripcion,
				abreviatura: dtoCampo.Abreviatura,
				id:          dtoCampo.Id,
			},
		}
	}
	return campo
}
