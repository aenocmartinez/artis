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

func (c *ColeccionDB) BuscarColeccionPorId(id int64) coleccion.Coleccion {
	return coleccion.Coleccion{}
}

func (c *ColeccionDB) ListarColecciones() []coleccion.DtoColeccion {
	return []coleccion.DtoColeccion{}
}

func (c *ColeccionDB) EliminarColeccion(id int64) bool {
	return false
}

func (c *ColeccionDB) ActualizarColeccion(coleccion coleccion.Coleccion) bool {
	return false
}

func (c *ColeccionDB) AgregarCampo(idColeccion int64, campo coleccion.CampoColeccion) bool {
	return false
}
