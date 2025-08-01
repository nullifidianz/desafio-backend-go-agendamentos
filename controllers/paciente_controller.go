package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/database"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/models"
)

func CadastrarPaciente(c *gin.Context) {
	var paciente models.Paciente

	if err := c.ShouldBindJSON(&paciente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	database.DB.Create(&paciente)
	c.JSON(http.StatusCreated, paciente)
}

func ListarPaciente(c *gin.Context) {
	var pacientes []models.Paciente
	database.DB.Find(&pacientes)
	c.JSON(http.StatusOK, pacientes)
}
