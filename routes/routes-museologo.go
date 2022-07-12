package routes

import (
	"artis/src/museologo/infraestructure/controller"

	"artis/src/museologo/infraestructure/repository/mysql"

	"github.com/gin-gonic/gin"
)

func CrearCampo(c *gin.Context) {
	ctr := controller.InstanceMuseologoController(mysql.InstanceCampoDB())
	err := ctr.CrearCampo(c)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}

func ListarCampos(c *gin.Context) {
	ctr := controller.InstanceMuseologoController(mysql.InstanceCampoDB())
	campos := ctr.ListarCampos(c)
	c.JSON(200, gin.H{"data": campos})
}

func VerCampo(c *gin.Context) {
	ctr := controller.InstanceMuseologoController(mysql.InstanceCampoDB())
	campo := ctr.VerCampo(c)
	c.JSON(200, gin.H{"data": campo})
}

func AgregarSubcampo(c *gin.Context) {
	ctr := controller.InstanceMuseologoController(mysql.InstanceCampoDB())
	err := ctr.AgregarSubcampo(c)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {

		c.JSON(200, gin.H{"message": "success"})
	}
}

func ListarSubcampos(c *gin.Context) {
	ctr := controller.InstanceMuseologoController(mysql.InstanceCampoDB())
	subcampos := ctr.ListarSubcampos(c)
	c.JSON(200, gin.H{"data": subcampos})
}

func EliminarCampo(c *gin.Context) {
	ctr := controller.InstanceMuseologoController(mysql.InstanceCampoDB())
	err := ctr.EliminarCampo(c)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}

func ActualizarCampo(c *gin.Context) {
	ctr := controller.InstanceMuseologoController(mysql.InstanceCampoDB())
	err := ctr.ActualizarCampo(c)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}

func QuitarCampo(c *gin.Context) {
	ctr := controller.InstanceMuseologoController(mysql.InstanceCampoDB())
	err := ctr.QuitarSubcampo(c)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}
