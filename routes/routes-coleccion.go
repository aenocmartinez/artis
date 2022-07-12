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
