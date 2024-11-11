package models

import (
	"github.com/goravel/framework/database/orm"
)

type Volunteer struct {
	orm.Model
	UserID   uint   `gorm:"not null;unique"`
	Training string `gorm:"type:varchar(100)"`  // entrenamiento completado, e.g., primeros auxilios
	Points   int    `gorm:"default:0"`
}
