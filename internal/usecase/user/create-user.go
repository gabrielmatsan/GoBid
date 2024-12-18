package user

import (
	"context"

	"github.com/gabrielmatsan/GoBid/internal/validator"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	// Check if the username is not blank
	eval.CheckField(validator.NotBlank(req.Username), "user_name", "This Field cannot be empty")

	// Check if the email is not blank
	eval.CheckField(validator.NotBlank(req.Email), "email", "This Field cannot be empty")

	// Check if the email is valid
	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "Invalid email address")

	// Check if the email is valid
	eval.CheckField(validator.NotBlank(req.Bio), "bio", "This Field cannot be empty")

	// Check if the bio is valid
	eval.CheckField(
		validator.MinChars(req.Bio, 10) &&
			validator.MaxChars(req.Bio, 255),
		"bio",
		"This Field must be between 10 and 255 characters")

	// Check if the password has at least 8 characters
	eval.CheckField(validator.MinChars(req.Password, 8), "password", "Password must be at least 8 characters")

	//
	return eval
}
