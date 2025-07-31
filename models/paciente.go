package models

import "gorm.io/gorm"

type Paciente struct {
	gorm.Model
	Nome     string `json:"nome"`
	Cpf      string `json:"cpf" gorm:"unique"`
	Telefone string `json:"telefone"`
}
