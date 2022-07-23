package routes

import (
	"artis/src/museologo/infraestructure/controller"
	"artis/src/museologo/infraestructure/repository/mysql"

	"github.com/gin-gonic/gin"
)

func CrearColeccion(c *gin.Context) {
	ctr := controller.InstanceColeccionController(mysql.InstanceColeccionDB(), mysql.InstanceCampoDB())
	err := ctr.CrearColeccion(c)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}

func ListarColecciones(c *gin.Context) {
	ctr := controller.InstanceColeccionController(mysql.InstanceColeccionDB(), mysql.InstanceCampoDB())
	c.JSON(200, gin.H{"data": ctr.ListarColecciones(c)})
}

func BuscarColeccion(c *gin.Context) {
	ctr := controller.InstanceColeccionController(mysql.InstanceColeccionDB(), mysql.InstanceCampoDB())
	coleccion := ctr.BuscarColeccion(c)
	if coleccion.Id == 0 {
		c.JSON(200, gin.H{"data": nil})
	} else {
		c.JSON(200, gin.H{"data": coleccion})
	}
}

func ActualizarColeccion(c *gin.Context) {
	ctr := controller.InstanceColeccionController(mysql.InstanceColeccionDB(), mysql.InstanceCampoDB())
	err := ctr.ActualizarColeccion(c)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}
