package rule

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	validation "github.com/nicklasjeppesen/going_internal/super/validation"
)

func PasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	fmt.Println("Passwrd Strength kaldt")
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	return hasUpper && hasLower && hasNumber
}

// RegisterAll custom validation rules
func RegisterAll() error {
	rules := map[string]func(fl validator.FieldLevel) bool{
		"password_strength": PasswordStrength,
	}

	for tag, fn := range rules {
		if err := validation.RegisterValidation(tag, fn); err != nil {
			fmt.Errorf("failed to register validation rule %q: %w", tag, err)
			panic(err)
		}
	}
	return nil
}
