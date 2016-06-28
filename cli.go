package main

import (
	"os"

	"github.com/james-nesbitt/wundertools-go/operation"
	"github.com/james-nesbitt/wundertools-go/config"
	"github.com/james-nesbitt/wundertools-go/log"
)

var (
	operationName string
	globalFlags map[string]string
	operationFlags []string

	app    config.Application
	logger log.Log
)

func init() {

	operationName, globalFlags, operationFlags = parseGlobalFlags(os.Args)

	// verbosity
	var verbosity int = log.VERBOSITY_MESSAGE
	if globalFlags["verbosity"] != "" {
		switch globalFlags["verbosity"] {
		case "message":
			verbosity = log.VERBOSITY_MESSAGE
		case "info":
			verbosity = log.VERBOSITY_INFO
		case "warning":
			verbosity = log.VERBOSITY_WARNING
		case "verbose":
			verbosity = log.VERBOSITY_DEBUG_LOTS
		case "debug":
			verbosity = log.VERBOSITY_DEBUG_WOAH
		case "staaap":
			verbosity = log.VERBOSITY_DEBUG_STAAAP
		}
	}
	logger = log.MakeCliLog("wundertools", os.Stdout, verbosity)

	workingDir, _ := os.Getwd()
	app = *config.DefaultApplication(workingDir)
}

func main() {

	if com, ok := operation.GetOperation(operationName); ok {

		com.Init(logger, &app)
		com.Execute(operationFlags...)

	} else {

		logger.Error("Unknown operation "+operationName)

	}

}
