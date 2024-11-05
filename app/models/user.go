package models

import (
	"errors"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	orm.Model
	Name     string
	Email    string
	Password string
	Avatar   string
}

// HashPassword encripta la contraseña y la asigna al usuario
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// UpdateFields actualiza los campos del usuario con los datos proporcionados
func (u *User) UpdateFields(name, email, password, avatar string) error {
	// Actualizar solo si los campos no están vacíos
	if name != "" {
		u.Name = name
	}
	if email != "" {
		u.Email = email
	}
	if avatar != "" {
		u.Avatar = avatar
	}
	if password != "" {
		if err := u.HashPassword(password); err != nil {
			return err
		}
	}

	// Guardar cambios en la base de datos
	result := facades.Orm().Query().Save(u)
	return result
}

func (u *User) SearchByEmail() error {
	return facades.Orm().Query().Where("email = ?", u.Email).FirstOrFail(&u)
}

// CheckPassword compara una contraseña en texto plano con el hash almacenado
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func validateUserUpdate(name, email, password, avatar string) error {
	if name == "" {
		return errors.New("El nombre es requerido")
	}
	if email == "" {
		return errors.New("El email es requerido")
	}
	// Puedes añadir más validaciones según sea necesario
	return nil
}

func (u *User) Update(name, email, password, avatar string) error {
	err := validateUserUpdate(name, email, password, avatar)
	if err != nil {
		return err
	}

	u.Name = name
	u.Email = email

	// Encriptar la nueva contraseña solo si se proporciona
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	// Asignar avatar si se proporciona
	if avatar != "" {
		u.Avatar = avatar
	}

	// Aquí agregas la lógica para guardar los cambios en la base de datos
	if err = facades.Orm().Query().Save(u); err != nil {
		return err
	}

	return nil
}
