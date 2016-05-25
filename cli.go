package main

import (
	"os"

	"github.com/james-nesbitt/wundertools-go/log"
	"github.com/james-nesbitt/wundertools-go/config"
	"github.com/james-nesbitt/wundertools-go/compose"
)

var (
	app config.Application
	logger log.Log
)

func init() {

	logger = log.MakeCliLog("wundertools", os.Stdout, log.VERBOSITY_MESSAGE)

	workingDir, _ := os.Getwd()
	app = *config.DefaultApplication(workingDir)

}

func main() {

	logger.Message("--SETTINGS--")
	logger.Debug(log.VERBOSITY_MESSAGE, "Name:", app.Name)
	logger.Debug(log.VERBOSITY_MESSAGE, "Author:", app.Author)
	logger.Debug(log.VERBOSITY_MESSAGE, "Environment:", app.Environment)

	logger.Message("--PATHS--")
	logger.Debug(log.VERBOSITY_MESSAGE, "Conf Path keys:", app.Paths.OrderedConfPathKeys())
	logger.Debug(log.VERBOSITY_MESSAGE, "All Paths:", app.Paths)

}
