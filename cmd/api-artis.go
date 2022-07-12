package main

import (
	"artis/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
	}

	srv := gin.New()
	srv.Use(gin.Logger(), gin.Recovery())

	srv.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	apiRoutes := srv.Group("/api-artis")
	apiRoutes.POST("/crear-campo", routes.CrearCampo)
	apiRoutes.GET("/listar-campos", routes.ListarCampos)
	apiRoutes.GET("/campo/:id", routes.VerCampo)
	apiRoutes.POST("/agregar-subcampo", routes.AgregarSubcampo)
	apiRoutes.GET("/campo/:id/subcampos", routes.ListarSubcampos)
	apiRoutes.DELETE("/eliminar-campo/:id", routes.EliminarCampo)
	apiRoutes.PUT("/actualizar-campo", routes.ActualizarCampo)
	apiRoutes.DELETE("/campo/:idCampo/subcampo/:idSubcampo", routes.QuitarCampo)

	srv.Run(":8080")
}
