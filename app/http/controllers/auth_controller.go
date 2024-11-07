package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/requests"
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
			"error": facades.Lang(ctx).Get("auth.login.email_not_found"),
		})
	}

	// Usar el metodo CheckPassword del modelo
	if !user.CheckPassword(password) {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": facades.Lang(ctx).Get("auth.login.failure"),
		})
	}

	token, err := facades.Auth(ctx).Login(&user)
	if err != nil {
		facades.Log().Error(err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": facades.Lang(ctx).Get("jwt.token_failure"),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": facades.Lang(ctx).Get("auth.login.success"),
		"token":   token,
	})
}

// Register function
func (c *AuthController) Register(ctx http.Context) http.Response {
	// Recibir datos de entrada
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")
	name := ctx.Request().Input("name")

	// Valida el formulario
	var registerUser requests.AuthRegisterRequest
	errors, err := ctx.Request().ValidateRequest(&registerUser)
	
	if err != nil {
		facades.Log().Debug(err)
	} else if errors != nil {
		facades.Log().Error(errors)
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": facades.Lang(ctx).Get("user.validator.form_invalid"),
		})
	}

	// Verificar si el email ya está registrado
	var existingUser models.User
	existingUser.Email = email
	err = existingUser.SearchByEmail()
	if err == nil {
		return ctx.Response().Json(http.StatusConflict, http.Json{
			"error": facades.Lang(ctx).Get("auth.create.email_already_registered"),
		})
	}

	// Crear y guardar el nuevo usuario
	newUser := models.User{
		Name:  name,
		Email: email,
	}
	err = newUser.HashPassword(password)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": facades.Lang(ctx).Get("auth.error_password"),
		})
	}

	if err := newUser.Create(); err != nil {
		facades.Log().Error(err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": facades.Lang(ctx).Get("user.create.failure"),
		})
	}
	// Generar un token JWT para el nuevo usuario (opcional)
	token, err := facades.Auth(ctx).Login(&newUser)
	if err != nil {
		facades.Log().Error(err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": facades.Lang(ctx).Get("jwt.token_failure"),
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
			"error": facades.Lang(ctx).Get("auth.logout.failure"),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": facades.Lang(ctx).Get("auth.logout.success"),
	})
}
