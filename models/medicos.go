package models

import "gorm.io/gorm"

type Medico struct {
	gorm.Model
	Nome          string `json:"nome"`
	Crm           string `json:"crm" gorm:"unique"`
	Especialidade string `json:"especialidade"`
}
