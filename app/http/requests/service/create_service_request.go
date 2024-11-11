package service

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type CreateServiceRequest struct {
	UserID      uint   `form:"user_id" json:"user_id"`
	ProtectorID uint   `form:"protector_id" json:"protector_id"`
	Type        string `form:"type" json:"type"`
	Location    string `form:"location" json:"location"`
	ScheduledAt string `form:"scheduled_at" json:"scheduled_at"`
}

func (r *CreateServiceRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateServiceRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"user_id":      "required|exists:users,id",
		"protector_id": "required_if:type,accompany,meet|exists:protectors,id",
		"type":         "required|in:urgency,accompany,meet,learn,warning",
		"location":     "required|json",
		"scheduled_at": "sometimes|date",
	}
}

func (r *CreateServiceRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"user_id.required":    facades.Lang(ctx).Get("service.user_id_required"),
		"protector_id.exists": facades.Lang(ctx).Get("service.protector_id_invalid"),
		"type.required":       facades.Lang(ctx).Get("service.type_required"),
		"location.required":   facades.Lang(ctx).Get("service.location_required"),
		"scheduled_at.date":   facades.Lang(ctx).Get("service.scheduled_at_invalid"),
	}
}

func (r *CreateServiceRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateServiceRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
