package main

import (
	"myapp/docs"
	"myapp/internal/app/jobs"
	provider "myapp/internal/app/providers/route"

	_ "github.com/lib/pq" // PostgreSQL driver
	app "github.com/nicklasjeppesen/going_internal/super/app"
)

func main() {

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
