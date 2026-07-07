package main

import (
	"fmt"
	"myapp/docs"
	"myapp/internal/app/jobs"
	provider "myapp/internal/app/providers/route"
	"regexp"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq" // PostgreSQL driver
	app "github.com/nicklasjeppesen/going_internal/super/app"
	validation "github.com/nicklasjeppesen/going_internal/super/validation"
)

// Should be removed, but here for the example
func passwordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	fmt.Println("Passwrd Strength kaldt")
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	return hasUpper && hasLower && hasNumber
}

func main() {

	// Registrer custom validation-regler FØR routes registreres,
	// så alle handlers har adgang til dem
	if err := validation.RegisterValidation("password_strength", passwordStrength); err != nil {
		panic(err)
	}

	// Create a new app instance
	var app = app.NewApp()

	// Register http and websocket routes
	provider.RegisterMaps(app.Router)

	// Register backgrounds job
	jobs.RegisterJobs(app.Scheduler)

	// Register swagger api - optional
	docs.RegisterSwagger(app.Router)

	// Start the application
	app.Start()
}
