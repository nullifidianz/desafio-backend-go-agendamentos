package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/controllers"
)

func Setup() *gin.Engine {
	routes := gin.Default()

	routes.POST("/medicos", controllers.CadastrarMedico)
	routes.GET("/medicos", controllers.ListarMedicos)

	routes.POST("/pacientes", controllers.CadastrarPaciente)
	routes.GET("/pacientes", controllers.ListarPaciente)

	routes.POST("/consultas", controllers.AgendarConsulta)
	routes.GET("/consultas", controllers.ListarConsultas)

	return routes
}
