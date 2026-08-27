package main

import (
	"myapp/docs"
	"myapp/internal/app/jobs"
	providerRoute "myapp/internal/app/providers/route"
	providerRule "myapp/internal/app/providers/rule"

	_ "github.com/lib/pq" // PostgreSQL driver
	app "github.com/nicklasjeppesen/going_internal/super/app"
)

func main() {

	// Create a new app instance
	var app = app.NewApp()

	// Register http and websocket routes
	providerRoute.RegisterMaps(app.Router)
	app.RegisterCustomRules(providerRule.RegisterCustomRule())

	// Register backgrounds job
	jobs.RegisterJobs(app.Scheduler)

	// Register swagger api - optional
	docs.RegisterSwagger(app.Router)

	// Start the application
	app.Start()
}
