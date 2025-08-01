package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/database"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("erro no arquivo .env")
	}
	database.Conectar()
	r := routes.Setup()
	r.Run(":8080")
}
