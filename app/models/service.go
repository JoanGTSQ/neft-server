package models

import (
	"github.com/goravel/framework/database/orm"
	"time"
)

type Service struct {
	orm.Model
	UserID        uint   `gorm:"not null"`                             // Usuario que solicita el servicio
	ProtectorID   uint   `gorm:"not null"`                             // Protector asignado
	Type          string `gorm:"type:varchar(50);not null"`            // urgencia, acompáñame, vamos a quedar
	Status        string `gorm:"type:varchar(50);default:'pendiente'"` // pendiente, en_proceso, finalizado
	StartTime     time.Time
	EndTime       time.Time
	LocationStart string `gorm:"type:geometry(Point,4326)"`
	LocationEnd   string `gorm:"type:geometry(Point,4326)"`
}
