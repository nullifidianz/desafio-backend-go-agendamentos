package models

import (
	"time"

	"gorm.io/gorm"
)

type Consulta struct {
	gorm.Model
	PacienteId uint      `json:"paciente_id"`
	MedicoId   uint      `json:"medico_id"`
	DataHora   time.Time `json:"data_hora"`
}
