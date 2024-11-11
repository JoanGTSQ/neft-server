package models

import (
	"github.com/goravel/framework/database/orm"
)

type Protector struct {
	orm.Model
	UserID       uint    `gorm:"not null;unique"`
	Rating       float64 `gorm:"type:decimal(3,2);default:0.0"`
	Status       string  `gorm:"type:varchar(50);default:'disponible'"` // disponible, ocupado, etc.
	LastLocation string  `gorm:"type:geometry(Point,4326)"`
}
