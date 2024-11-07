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
	
	if user.ID == targetUser.ID {
		return access.NewAllowResponse()
	} else {
		return access.NewDenyResponse(facades.Lang(ctx).Get("user.policies.unmatch_users"))
	}
}
