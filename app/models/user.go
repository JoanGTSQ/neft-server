package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
)

type User struct {
	orm.Model
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Role     string `json:"role" gorm:"column:role"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
}

// HashPassword encripta la contraseña y la asigna al usuario
func (u *User) HashPassword(password string) error {
	hashedPassword, err := facades.Hash().Make(password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

func (u *User) Create() error {
	err := facades.Orm().Query().Create(&u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) SearchByEmail() error {
	return facades.Orm().Query().Where("email = ?", u.Email).FirstOrFail(&u)
}

// Método para obtener un usuario por su email
func UserByEmail(email string) (*User, error) {
	var user User
	if err := facades.Orm().Query().Where("email = ?", email).FirstOrFail(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) SearchById() error {
	return facades.Orm().Query().Where("id = ?", u.ID).FindOrFail(&u)
}

// CheckPassword compara una contraseña en texto plano con el hash almacenado
func (u *User) CheckPassword(password string) bool {
	if facades.Hash().Check(password, u.Password) {
		return true
	}
	return false
}

func (u *User) Update() error {
	return facades.Orm().Query().Save(u)
}
