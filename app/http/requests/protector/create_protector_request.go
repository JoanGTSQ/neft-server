package protector

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type CreateProtectorRequest struct {
	UserID   uint    `form:"user_id" json:"user_id"`
	Status   string  `form:"status" json:"status"`
	Rating   float64 `form:"rating" json:"rating"`
	Location string  `form:"location" json:"location"` // Para almacenar la ubicación en JSON
}

func (r *CreateProtectorRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateProtectorRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"user_id":  "required|exists:users,id",
		"status":   "required|in:available,busy,unavailable",
		"rating":   "sometimes|numeric|min:0.0|max:5.0",
		"location": "required|json",
	}
}

func (r *CreateProtectorRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"user_id.required":  facades.Lang(ctx).Get("protector.user_id_required"),
		"user_id.exists":    facades.Lang(ctx).Get("protector.user_id_invalid"),
		"status.required":   facades.Lang(ctx).Get("protector.status_required"),
		"status.in":         facades.Lang(ctx).Get("protector.status_invalid"),
		"rating.numeric":    facades.Lang(ctx).Get("protector.rating_invalid"),
		"rating.min":        facades.Lang(ctx).Get("protector.rating_min"),
		"rating.max":        facades.Lang(ctx).Get("protector.rating_max"),
		"location.required": facades.Lang(ctx).Get("protector.location_required"),
		"location.json":     facades.Lang(ctx).Get("protector.location_invalid"),
	}
}

func (r *CreateProtectorRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateProtectorRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}