package main

import (
	"context"
	"os"
	"time"
	"tsf-cron"
	"tsf-cron/pkg/core/log"
	"tsf-cron/pkg/core/log/zerolog"

	"go.temporal.io/sdk/client"
)

func main() {
	logDir := "./tmp/log/start/"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fileName := logDir + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	log.New(zerolog.New(log.WithOutput(logFile)))
	log.Infof("hhh")
	log.Errorf("start")

	c, err := client.Dial(client.Options{
		HostPort: app.Addr,
	})
	if err != nil {
		log.Fatalf("unable to create Temporal client: %s", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:           "greeting-workflow",
		TaskQueue:    app.GreetingTaskQueue,
		CronSchedule: "* * * * *",
	}

	name := "world"
	_, err = c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflow, name)
	if err != nil {
		log.Fatalf("unable to complete Workflow: %s", err)
	}
}
