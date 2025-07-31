package database

import (
	"log"

	"github.com/nullifidianz/desafio-backend-go-agendamentos/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conectar() {
	var err error
	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro com o db: ", err)
	}

	DB.AutoMigrate(&models.Medico{}, &models.Paciente{}, &models.Consulta{})
}
