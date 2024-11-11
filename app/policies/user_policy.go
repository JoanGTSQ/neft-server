package policies

import (
	"context"
	"github.com/goravel/framework/auth/access"
	contractsaccess "github.com/goravel/framework/contracts/auth/access"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type UserPolicy struct {
}

func NewUserPolicy() *UserPolicy {
	return &UserPolicy{}
}

func (r *UserPolicy) Update(ctx context.Context, arguments map[string]any) contractsaccess.Response {
	user := ctx.Value("user").(models.User)
	targetUser := arguments["targetUser"].(models.User)
	// Verificar si el correo electrónico ya está registrado
	existingUser, err := models.UserByEmail(targetUser.Email)
	if err != nil {
		facades.Log().Error(err)
		return access.NewDenyResponse(facades.Lang(ctx).Get("user.policies.email_check_error"))
	}

	if user.Role == "admin" {
		return access.NewAllowResponse()
	}

	// Si el correo electrónico ya está registrado y no es del usuario actual
	if existingUser != nil && existingUser.ID != user.ID {
		// El correo ya pertenece a otro usuario
		return access.NewDenyResponse(facades.Lang(ctx).Get("user.policies.email_already_taken"))
	}

	// Si el usuario está intentando actualizar su propia información, se permite
	if user.ID == targetUser.ID {
		return access.NewAllowResponse()
	}

	// Si no es el mismo usuario, denegar
	return access.NewDenyResponse(facades.Lang(ctx).Get("user.policies.unmatch_users"))
}
