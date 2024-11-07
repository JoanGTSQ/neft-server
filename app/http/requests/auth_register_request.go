package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AuthRegisterRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *AuthRegisterRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AuthRegisterRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		// The keys are consistent with the incoming keys.
		"name":     "required|max_len:255",
		"email":    "required|max_len:255",
		"password": "required|max_len:255",
	}
}

func (r *AuthRegisterRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthRegisterRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthRegisterRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
