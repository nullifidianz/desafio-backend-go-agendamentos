package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/auth"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/database"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/models"
	"golang.org/x/crypto/bcrypt"
)

func RegistrarUsuario(c *gin.Context) {
	var input models.Usuario
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Senha), bcrypt.DefaultCost)
	input.Senha = string(hash)

	database.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

func Login(c *gin.Context) {
	var input struct {
		Nome  string `json:"nome"`
		Email string `json:"email"`
		Senha string `json:"senha"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var usuario models.Usuario
	if err := database.DB.Where("email = ?", input.Email).First(&usuario).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Credenciais inválidas"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(input.Senha)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Credenciais inválidas"})
		return
	}

	token, _ := auth.GerarToken(usuario.ID, usuario.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
