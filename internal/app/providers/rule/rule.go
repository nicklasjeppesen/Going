package rule

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
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
func RegisterCustomRule() map[string]func(fl validator.FieldLevel) bool {
	return map[string]func(fl validator.FieldLevel) bool{
		// register our custom rules.
		"password_strength": PasswordStrength,
	}
}
