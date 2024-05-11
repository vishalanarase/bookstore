package main

import (
	"path"
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/api"
	"github.com/vishalanarase/bookstore/internal/config"
)

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			return "", fileName
		},
	})
}

func main() {
	log.Info("It's API")

	envConfig := config.Config("../../")

	app := api.NewApplication()
	app.Start(envConfig)
}
