package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}
func (u *UserController) Update(ctx http.Context) http.Response {
	// Obtener el usuario del contexto
	userValue := ctx.Value("user")
	if userValue == nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": "Usuario no encontrado en el contexto",
		})
	}

	user, ok := userValue.(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Error al convertir el usuario.",
		})
	}

	// Obtener datos del request
	id := ctx.Request().Input("ID") // ID del usuario a actualizar
	name := ctx.Request().Input("name")
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")
	avatar := ctx.Request().Input("avatar")

	if facades.Gate().WithContext(ctx).Denies("update-user", map[string]any{
		"targetUserId": id,
	}) {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": "No puedes modificar este usuario",
		})
	}

	// Validar los datos
	if err := user.Update(name, email, password, avatar); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Datos actualizados con éxito",
		"user":    user,
	})
}
