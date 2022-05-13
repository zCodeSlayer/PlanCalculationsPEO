package validators

import (
	"go-postgres/models"
	"gopkg.in/go-playground/validator.v9"
)

func UserValidation(user models.User) error {
	return validator.New().Struct(user)
}

func GroupValidation(group models.Group) error {
	return validator.New().Struct(group)
}
