package mysql

import (
	"artis/config/database"
	"artis/src/museologo/domain/coleccion"
	"bytes"
	"database/sql"
	"log"
)

type ColeccionDB struct {
	conn *sql.DB
}

func InstanceColeccionDB() *ColeccionDB {
	return &ColeccionDB{
		conn: database.Instance().Connection(),
	}
}

func (m *ColeccionDB) Crear(coleccion coleccion.Coleccion) bool {
	var query bytes.Buffer
	query.WriteString("INSERT INTO colecciones(nombre) VALUES (?)")
	stmt, err := m.conn.Prepare(query.String())

	if err != nil {
		log.Println(err.Error())
		return false
	}

	result, err := stmt.Exec(coleccion.Nombre)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	coleccion.Id = id

	defer stmt.Close()
	return true
}

func (m *ColeccionDB) BuscarColeccionPorId(id int64) coleccion.Coleccion {
	var oColeccion coleccion.Coleccion
	var nombre string
	var idColeccion int64
	var query bytes.Buffer
	query.WriteString("SELECT id, nombre FROM colecciones WHERE id = ?")

	row := m.conn.QueryRow(query.String(), id)
	row.Scan(&idColeccion, &nombre)

	oColeccion = coleccion.Coleccion{
		Id:     idColeccion,
		Nombre: nombre,
	}

	return oColeccion
}

func (m *ColeccionDB) ListarColecciones() []coleccion.DtoColeccion {
	var listaColecciones []coleccion.DtoColeccion
	var query bytes.Buffer

	query.WriteString("SELECT id, nombre FROM colecciones ORDER BY nombre ")

	rows, err := m.conn.Query(query.String())
	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var nombre string

		rows.Scan(&id, &nombre)

		listaColecciones = append(listaColecciones, coleccion.DtoColeccion{
			Id:     id,
			Nombre: nombre,
		})
	}
	return listaColecciones
}

func (c *ColeccionDB) EliminarColeccion(id int64) bool {
	return false
}

func (c *ColeccionDB) ActualizarColeccion(coleccion coleccion.Coleccion) bool {
	var query bytes.Buffer
	query.WriteString("UPDATE colecciones SET nombre=?, updated_at=NOW() WHERE id = ?")
	stmt, err := c.conn.Prepare(query.String())

	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = stmt.Exec(coleccion.Nombre, coleccion.Id)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	defer stmt.Close()
	return true
}

func (c *ColeccionDB) AgregarCampo(idColeccion int64, campo coleccion.CampoColeccion) bool {
	return false
}
