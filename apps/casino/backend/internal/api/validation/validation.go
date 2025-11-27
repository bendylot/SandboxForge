package validation

import (
	"strings"

	"casino.sandboxforge.tech/casino-api/internal/api/dto"
	"github.com/go-playground/validator/v10"
)

var v = validator.New()

func NormalizeRegisterReq(r *dto.RegisterReq) {
	r.Login = strings.ToLower(strings.TrimSpace(r.Login))
	r.Email = strings.ToLower(strings.TrimSpace(r.Email))
}

func ValidateRegisterReq(r *dto.RegisterReq) error {
	return v.Struct(r)
}

func IsValidationError(err error) bool {
	_, ok := err.(validator.ValidationErrors)
	return ok
}