package user

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateUserRequest struct {
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
}

func (r *UpdateUserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateUserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		// The keys are consistent with the incoming keys.
		"name":  "required|max_len:255",
		"email": "required|max_len:255",
	}
}

func (r *UpdateUserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateUserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateUserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
