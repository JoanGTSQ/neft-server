package policies

import (
	"context"
	"goravel/app/models"
	"strconv"

	"github.com/goravel/framework/auth/access"
	contractsaccess "github.com/goravel/framework/contracts/auth/access"
)

type UserPolicy struct {
}

func NewUserPolicy() *UserPolicy {
	return &UserPolicy{}
}

func (r *UserPolicy) Update(ctx context.Context, arguments map[string]any) contractsaccess.Response {
	user := ctx.Value("user").(models.User)
	targetUserID := arguments["targetUserId"]
	userID := strconv.FormatUint(uint64(user.ID), 10)

	if userID == targetUserID {
		return access.NewAllowResponse()
	} else {
		return access.NewDenyResponse("You do not own this post.")
	}
}
