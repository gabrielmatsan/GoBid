package user

import (
	"context"

	"github.com/gabrielmatsan/GoBid/internal/validator"
)

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginUserRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "Invalid email")
	eval.CheckField(validator.NotBlank(req.Password), "password", "Password cannot be blank")

	return eval
}
