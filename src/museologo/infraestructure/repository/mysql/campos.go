package mysql

import (
	"artis/config/database"
	"artis/src/museologo/domain/campos"
	"bytes"
	"database/sql"
	"log"
)

type CampoDB struct {
	conn *sql.DB
}

func InstanceCampoDB() *CampoDB {
	return &CampoDB{
		conn: database.Instance().Connection(),
	}
}

func (m *CampoDB) CrearCampo(campo campos.InterfaceCampo) bool {
	var query bytes.Buffer
	query.WriteString("INSERT INTO campos(nombre, descripcion, abreviatura) VALUES (?, ?, ?)")

	stmt, err := m.conn.Prepare(query.String())

	if err != nil {
		log.Println(err.Error())
		return false
	}

	result, err := stmt.Exec(campo.Nombre(), campo.Descripcion(), campo.Abreviatura())
	if err != nil {
		log.Println(err.Error())
		return false
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	campo.SetId(id)

	defer stmt.Close()
	return true
}

func (m *CampoDB) EliminarCampo(id int64) bool {
	var query bytes.Buffer
	query.WriteString("DELETE FROM campos WHERE id = ?")

	stmt, err := m.conn.Prepare(query.String())

	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	defer stmt.Close()
	return true
}

func (m *CampoDB) ActualizarCampo(campo *campos.Campo) bool {
	var query bytes.Buffer
	query.WriteString("UPDATE campos SET nombre=?, descripcion=?, abreviatura=?, updated_at=NOW() WHERE id = ?")

	stmt, err := m.conn.Prepare(query.String())

	if err != nil {
		log.Println(err.Error())
		return false
	}

	result, err := stmt.Exec(campo.Nombre(), campo.Descripcion(), campo.Abreviatura(), campo.Id())
	if err != nil {
		log.Println(err.Error())
		return false
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	campo.SetId(id)

	defer stmt.Close()
	return true
}

func (m *CampoDB) ListarCampos() []campos.DtoCampo {
	var listaCampos []campos.DtoCampo
	var query bytes.Buffer

	query.WriteString("SELECT ")
	query.WriteString("distinct campos.id, campos.nombre, campos.descripcion, campos.abreviatura, IF (subcampos.campo_id is null, 0, 1) as esCompuesto ")
	query.WriteString("FROM ")
	query.WriteString("campos ")
	query.WriteString("left join subcampos on subcampos.campo_id = campos.id ")
	query.WriteString("ORDER BY campos.nombre ")

	rows, err := m.conn.Query(query.String())
	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var nombre, descripcion, abreviatura string
		var esCompuesto bool

		rows.Scan(&id, &nombre, &descripcion, &abreviatura, &esCompuesto)

		listaCampos = append(listaCampos, campos.DtoCampo{
			Id:          id,
			Nombre:      nombre,
			Abreviatura: abreviatura,
			Descripcion: descripcion,
			EsCompuesto: esCompuesto,
		})
	}
	return listaCampos
}

func (m *CampoDB) BuscarCampoPorNombre(nombre string) campos.InterfaceCampo {
	return nil
}

func (m *CampoDB) BuscarCampoPorId(id int64) campos.InterfaceCampo {
	var nombre, descripcion, abreviatura string
	var esCompuesto bool
	var campo campos.InterfaceCampo

	var query bytes.Buffer
	query.WriteString("SELECT ")
	query.WriteString("distinct campos.id, campos.nombre, campos.descripcion, campos.abreviatura, IF (subcampos.campo_id is null, 0, 1) as esCompuesto ")
	query.WriteString("FROM ")
	query.WriteString("campos ")
	query.WriteString("left join subcampos on subcampos.campo_id = campos.id ")
	query.WriteString("WHERE campos.id = ? ")

	row := m.conn.QueryRow(query.String(), id)
	row.Scan(&id, &nombre, &descripcion, &abreviatura, &esCompuesto)

	campo = campos.InstanceCampoSimple()
	if esCompuesto {
		campo = campos.InstanceCampoCompuesto()
	}
	campo.SetId(id)
	campo.SetNombre(nombre)
	campo.SetAbreviatura(abreviatura)
	campo.SetDescripcion(descripcion)

	return campo
}

func (m *CampoDB) AgregarSubcampo(idCampo, idSubcampo, orden int64) bool {
	var query bytes.Buffer
	query.WriteString("INSERT INTO subcampos(campo_id, subcampo_id, orden) VALUES (?, ?, ?)")

	stmt, err := m.conn.Prepare(query.String())

	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = stmt.Exec(idCampo, idSubcampo, orden)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	defer stmt.Close()
	return true
}

func (m *CampoDB) QuitarSubcampo(idCampo, idSubcampo int64) bool {
	var query bytes.Buffer
	query.WriteString("DELETE FROM subcampos WHERE campo_id = ? AND subcampo_id = ?")

	stmt, err := m.conn.Prepare(query.String())

	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = stmt.Exec(idCampo, idSubcampo)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	defer stmt.Close()
	return true
}

func (m *CampoDB) ListarSubcampos(idCampo int64) []campos.DtoCampo {
	var subcampos []campos.DtoCampo
	var query bytes.Buffer
	query.WriteString("SELECT ")
	query.WriteString("c2.id, c2.nombre, c2.descripcion, c2.abreviatura ")
	query.WriteString("FROM campos c1 ")
	query.WriteString("INNER JOIN subcampos sc on sc.campo_id = c1.id ")
	query.WriteString("INNER JOIN campos c2 on c2.id = sc.subcampo_id ")
	query.WriteString("WHERE c1.id = ? and c1.id <> c2.id ")
	query.WriteString("ORDER BY sc.orden ")

	rows, err := m.conn.Query(query.String(), idCampo)
	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var nombre, descripcion, abreviatura string

		rows.Scan(&id, &nombre, &descripcion, &abreviatura)

		subcampos = append(subcampos, campos.DtoCampo{
			Id:          id,
			Nombre:      nombre,
			Abreviatura: abreviatura,
			Descripcion: descripcion,
		})
	}
	return subcampos
}
