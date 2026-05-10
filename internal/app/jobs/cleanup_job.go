package jobs

import (
	"log"
	"time"

	"github.com/nicklasjeppesen/going_internal/super/channels"
	"github.com/nicklasjeppesen/going_internal/super/jobs"
)

func CleanupJob() Job {
	return jobs.Job{
		Title:          "CleanJob",
		Interval:       5 * time.Second,
		Runner:         Runner,
		TerminateAfter: 10 * time.Second,
	}
}

func Runner(job Job) {
	log.Println("cleanup: start")
	for i := 0; i < 20; i++ {
		if job.IsInterrupted() {
			HandleIfJobInterrupted(job)
			return
		}
		time.Sleep(1 * time.Second)
		log.Printf("cleanup: chunk %d færdig", i+1)
		job.SendMessageToSocket(channels.Socket{URL: "/ws/message", Message: "Message to socket"})
	}
	log.Println("CLEANUP: Job finised")
}

func HandleIfJobInterrupted(job Job) {
	if job.InterruptedByDeadlineExceeded() {
		log.Println("OH no, exceed Deadline limit")
	} else if job.InterruptedByCanceled() {
		log.Println("cleanup: Stop by ctr c command")
	}
}
