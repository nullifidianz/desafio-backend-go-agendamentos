package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/controllers"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/middleware"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/registrar", controllers.RegistrarUsuario)

	autorizadas := r.Group("/")
	autorizadas.Use(middleware.AuthMiddleware())

	autorizadas.GET("/consultas", middleware.Autorizar("medico", "admin"), controllers.ListarConsultas)

	autorizadas.POST("/medicos", middleware.Autorizar("admin"), controllers.CadastrarMedico)

	return r
}
