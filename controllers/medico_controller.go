package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/database"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/models"
)

func CadastrarMedico(c *gin.Context) {
	var medico models.Medico
	if err := c.ShouldBindJSON(&medico); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	database.DB.Create(&medico)
	c.JSON(http.StatusCreated, medico)
}

func ListarMedicos(c *gin.Context) {
	var medicos []models.Medico
	database.DB.Find(&medicos)
	c.JSON(http.StatusOK, medicos)
}
