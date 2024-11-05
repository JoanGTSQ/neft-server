package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"golang.org/x/crypto/bcrypt"
	"goravel/app/models"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// Login permite al usuario autenticarse
func (a *AuthController) Login(ctx http.Context) http.Response {
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")

	var user models.User
	user.Email = email
	err := user.SearchByEmail()
	if err != nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": "Usuario no encontrado o credenciales incorrectas",
		})
	}

	// Usar el metodo CheckPassword del modelo
	if !user.CheckPassword(password) {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": "Credenciales incorrectas",
		})
	}

	token, err := facades.Auth(ctx).Login(&user)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Error al generar el token de autenticación",
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Inicio de sesión exitoso",
		"token":   token,
	})
}

// Register function
func (c *AuthController) Register(ctx http.Context) http.Response {
	// Recibir datos de entrada
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")
	name := ctx.Request().Input("name")

	// Validar datos de entrada
	if email == "" || password == "" || name == "" {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": "Email, nombre y contraseña son obligatorios",
		})
	}

	// Verificar si el email ya está registrado
	var existingUser models.User
	existingUser.Email = email
	err := existingUser.SearchByEmail()
	if err == nil {
		return ctx.Response().Json(http.StatusConflict, http.Json{
			"error": "El email ya está registrado",
		})
	}

	// Encriptar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Error al encriptar la contraseña",
		})
	}

	// Crear y guardar el nuevo usuario
	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	result := facades.Orm().Query().Create(&newUser)
	if result != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Error al crear el usuario",
		})
	}

	// Generar un token JWT para el nuevo usuario (opcional)
	token, err := facades.Auth(ctx).Login(&newUser)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "No se pudo generar el token",
		})
	}

	// Responder con el token JWT
	return ctx.Response().Json(http.StatusOK, http.Json{
		"token": token,
	})
}

func (c *AuthController) Logout(ctx http.Context) http.Response {
	err := facades.Auth(ctx).Logout()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "No se pudo cerrar sesión",
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Sesión cerrada correctamente",
	})
}
