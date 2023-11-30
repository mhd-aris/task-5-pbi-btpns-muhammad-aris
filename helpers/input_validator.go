package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/app"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/models"
)

var UserValidator *validator.Validate

func init() {
	UserValidator = validator.New()
}

func ValidateUser(user models.User) error {
	return UserValidator.Struct(user)
}
func ValidateLoginInput(input app.LoginInput) error {
	return UserValidator.Struct(input)
}