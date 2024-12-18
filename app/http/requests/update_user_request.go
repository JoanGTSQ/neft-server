package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type UpdateUserRequest struct {
	ID       uint   `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Avatar   string `form:"avatar" json:"avatar"`
	Password string `form:"password" json:"password"`
	Role     string `form:"role" json:"role"`
}

func (r *UpdateUserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateUserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		// The keys are consistent with the incoming keys.
		"id":       "required|numeric",
		"name":     "minLen:3|maxLen:20",
		"password": "password|minLen:8|maxLen:20",
		"email":    "email",
		"avatar":   "url",
		"role":     "in:admin,user,guest",
	}
}

func (r *UpdateUserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"id.required":       facades.Lang(ctx).Get("user.validator.id_required"),
		"id.numeric":        facades.Lang(ctx).Get("user.validator.id_invalid"),
		"password.required": facades.Lang(ctx).Get("user.validator.password_required"),
		"password.password": facades.Lang(ctx).Get("user.validator.password_invalid"),
		"password.min":      facades.Lang(ctx).Get("user.validator.password_min"),
		"password.max":      facades.Lang(ctx).Get("user.validator.password_max"),
		"name.required":     facades.Lang(ctx).Get("user.validator.name_required"),
		"email.required":    facades.Lang(ctx).Get("user.validator.email_required"),
		"avatar.required":   facades.Lang(ctx).Get("user.validator.avatar_required"),
		"role.required":     facades.Lang(ctx).Get("user.validator.role_required"),
		"name.min":          facades.Lang(ctx).Get("user.validator.name_min"),
		"name.max":          facades.Lang(ctx).Get("user.validator.name_max"),
		"email.email":       facades.Lang(ctx).Get("user.validator.email_invalid"),
		"avatar.url":        facades.Lang(ctx).Get("user.validator.avatar_invalid"),
		"role.in":           facades.Lang(ctx).Get("user.validator.role_invalid"),
	}
}

func (r *UpdateUserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateUserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
