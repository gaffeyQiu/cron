package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"tsf-cron"
	"tsf-cron/config"
	"tsf-cron/internal/initialize"
	"tsf-cron/pkg/core/log"
	"tsf-cron/pkg/core/log/zerolog"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

var (
	configPath string
)

// 解析命令行参数
func init() {
	flag.StringVar(&configPath, "conf", "", "config file path, example: -conf config.yml")
}

func main() {
	flag.Parse()
	Init()

	temporalServer := fmt.Sprintf("%s:%d", config.GetString("TemporalServer.Ip"), config.GetInt("TemporalServer.Port"))
	log.Infof("temporal server: %s", temporalServer)

	c, err := client.Dial(client.Options{
		HostPort: temporalServer,
		Logger:   log.GetLogger(),
	})
	if err != nil {
		log.Fatalf("unable to create Temporal client: %s", err)
	}
	defer c.Close()

	w := worker.New(c, config.GetString("Worker.TaskQueue"), worker.Options{})
	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.ComposeGreeting)

	log.Info("Woker Running")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("unable to start Worker: %s", err)
	}
}

func Init() {
	config.Init(configPath)

	// 创建程序日志目录
	logDir := fmt.Sprintf("%s/worker/", config.GetString("App.Log.Dir"))
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fileName := logDir + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// 初始化 log
	log.New(zerolog.New(log.WithOutput(logFile)))
	log.Info("Success Init Log Config...")

	initialize.InitGlobal()
	log.Info("Success Init Global ...")
}
