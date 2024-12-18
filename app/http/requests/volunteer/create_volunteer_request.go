package volunteer

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type CreateVolunteerRequest struct {
	UserID            uint `form:"user_id" json:"user_id"`
	TrainingCompleted bool `form:"training_completed" json:"training_completed"`
	Points            int  `form:"points" json:"points"`
}

func (r *CreateVolunteerRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateVolunteerRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"user_id":            "required|exists:users,id",
		"training_completed": "required|boolean",
		"points":             "required|integer|min:0",
	}
}

func (r *CreateVolunteerRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"user_id.required":            facades.Lang(ctx).Get("volunteer.user_id.required"),
		"training_completed.required": facades.Lang(ctx).Get("volunteer.training.required"),
		"training_completed.boolean":  facades.Lang(ctx).Get("volunteer.training.boolean"),
		"points.required":             facades.Lang(ctx).Get("volunteer.points.required"),
		"pints.integer":               facades.Lang(ctx).Get("volunteer.points.integer"),
		"points.min":                  facades.Lang(ctx).Get("volunteer.points.min"),
	}
}

func (r *CreateVolunteerRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateVolunteerRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
