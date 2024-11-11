package models

import (
	"github.com/goravel/framework/database/orm"
	"time"
)

type Warning struct {
	orm.Model
	ReportedBy  uint   `gorm:"not null"` // Usuario que reporta la advertencia
	Description string `gorm:"type:text;not null"`
	Location    string `gorm:"type:geometry(Point,4326)"`
	CreatedAt   time.Time
}
