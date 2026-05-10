package jobs

import (
	jobs "github.com/nicklasjeppesen/going_internal/super/jobs"
)

type Job = jobs.Job

//--------------------------------------------------------------------------
// 								Base
//--------------------------------------------------------------------------
//
// Here is where you can register Period jobs for your application.
// These jobs are started when you start your web server.
// If the server receive a shutdown, The job has 30 second to stop it task.
//
// The job will also stop if a timeoutdeadline is set in the job
//
// If you job run longer than the next cycle, the app will not start
// a new instance of the job, but allowed your job finished running
//
//

// Register.
//
// # Register all periodsjobs for the app
//
// param s: *jobs.scheduler:
//
// Exemple:
func RegisterJobs(schedule *jobs.Scheduler) {
	//schedule.CreateJob(CleanupJob())
}
