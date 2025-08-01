package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/database"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/models"
)

func AgendarConsulta(c *gin.Context) {
	var consulta models.Consulta
	if err := c.ShouldBindJSON(&consulta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	var conflito models.Consulta
	agendamentoDuplo := database.DB.Where("medico_id = ? AND data_hora = ?", consulta.MedicoId, consulta.DataHora).First(&conflito)
	if agendamentoDuplo.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"erro": "conflito na data e hora da consulta"})
		return
	}

	database.DB.Create(&consulta)
	c.JSON(http.StatusCreated, consulta)
}

func ListarConsultas(c *gin.Context) {
	var consultas []models.Consulta
	database.DB.Find(&consultas)
	c.JSON(http.StatusOK, consultas)
}
