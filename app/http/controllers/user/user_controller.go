package user

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/requests"
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

func (u *UserController) Show(ctx http.Context) http.Response {
	user := ctx.Value("user").(models.User)
	return ctx.Response().Json(http.StatusOK, user)
}

func (u *UserController) Update(ctx http.Context) http.Response {
	// Obtener el usuario del contexto
	user := ctx.Value("user").(models.User)

	var updateUser requests.UpdateUserRequest
	errors, err := ctx.Request().ValidateRequest(&updateUser)

	if err != nil {
		facades.Log().Error(err)
	} else if errors != nil {
		facades.Log().Error(errors)
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": facades.Lang(ctx).Get("user.validator.form_invalid"),
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
			"error": facades.Lang(ctx).Get("user.policies.denied"),
		})
	}

	// Validar los datos
	if err := user.Update(name, email, password, avatar); err != nil {
		facades.Log().Error(err)
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": facades.Lang(ctx).Get("user.update.success"),
		"user":    user,
	})
}
