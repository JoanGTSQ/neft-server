package user

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/database/orm"
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
	} else if errors != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": errors.All(),
		})
	}

	// Cogemos el usuario del formulario
	userToUpdate := models.User{
		Name:   updateUserRequest.Name,
		Email:  updateUserRequest.Email,
		Avatar: updateUserRequest.Avatar,
		Role:   updateUserRequest.Role,
		Model: orm.Model{
			ID: updateUserRequest.ID,
		},
	}
	// Verificamos las politicas
	response := facades.Gate().WithContext(ctx).Inspect("update-user", map[string]any{
		"targetUser": userToUpdate,
	})
	if !response.Allowed() {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": response.Message(),
		})
	}

	// Actualizamos el usuario
	if err := userToUpdate.Update(); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": facades.Lang(ctx).Get("user.update.failure"),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": facades.Lang(ctx).Get("user.update.success"),
		"user":    userToUpdate,
	})
}
