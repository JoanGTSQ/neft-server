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
	// Validamos el formulario
	var updateUserRequest requests.UpdateUserRequest
	errors, err := ctx.Request().ValidateRequest(&updateUserRequest)

	if err != nil {
		facades.Log().Error(err)
	} else if errors != nil {
		facades.Log().Error(errors)
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": facades.Lang(ctx).Get("user.validator.form_invalid"),
		})
	}

	// Cogemos el usuario del formulario
	var userToUpdate models.User
	ctx.Request().Bind(&userToUpdate)
	
	// Verificamos las politicas
	if facades.Gate().WithContext(ctx).Denies("update-user", map[string]any{
		"targetUser": userToUpdate,
	}) {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": facades.Lang(ctx).Get("user.policies.denied"),
		})
	}

	// Actualizamos el usuario
	if err := userToUpdate.Update(); err != nil {
		facades.Log().Error(err)
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": facades.Lang(ctx).Get("user.update.success"),
		"user":    userToUpdate,
	})
}
