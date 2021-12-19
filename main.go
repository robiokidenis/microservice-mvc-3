package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/robiokidenis/microservice-mvc-2/conf"
	"github.com/robiokidenis/microservice-mvc-3/jobs/payment"
	"time"
)

// how to make even scheduler with golang
func main() {
	schedule := gocron.NewScheduler()
	schedule.Every(1).Minutes().Do(task)
	<-schedule.Start()
}

func task() {
	cfg := conf.NewConfig()

	taskList := cfg.GetStrings("task")

	for _, tasks := range taskList {
		now := time.Now().Format("15:04")
		executingTime := cfg.GetString(tasks + ".executing_time")

		if now == executingTime {
			fmt.Printf("Running task %s", cfg.GetString(tasks+".description"))
			chooseTask(tasks)
		}
	}
}

func chooseTask(task string) {
	switch task {
	case "update_expired_payment":
		payment.UpdateExpired()
	case "jobs_2":
		// TO DO call some function
	case "jobs_3":
		// TO DO call some function
	}
}
