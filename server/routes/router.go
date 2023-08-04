package routes

import (
	"bitbucket.org/adson14/api_golang_ii.git/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		imoveis := main.Group("imoveis")
		{
			imoveis.GET("/", controllers.MostrarImoveis)
			imoveis.GET("/:id", controllers.MostrarImovel)
			imoveis.POST("/", controllers.CriarImovel)
			imoveis.PUT("/", controllers.AtualizarImovel)
			imoveis.DELETE("/:id", controllers.RemoveImovel)

		}

		proprietarios := main.Group("proprietarios")
		{
			proprietarios.GET("/", controllers.MostrarProprietarios)
			proprietarios.GET("/:id", controllers.MostrarProprietario)
			proprietarios.POST("/", controllers.CriarProprietario)
			proprietarios.PUT("/", controllers.AtualizarProprietario)
			proprietarios.DELETE("/:id", controllers.RemoveProprietario)

		}
	}

	return router
}
