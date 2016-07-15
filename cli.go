package main

/**
 * @TODO This will likely soon be replaced with github.com/urfave/cli
 */

import (
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/config"
	"github.com/james-nesbitt/wundertools-go/operation"
)

var (
	operationName  string
	globalFlags    map[string]string
	operationFlags []string

	app    config.Application
)

func init() {

	operationName, globalFlags, operationFlags = parseGlobalFlags(os.Args)

	// verbosity
	if globalFlags["verbosity"] != "" {
		if level, err := log.ParseLevel(globalFlags["verbosity"]); err==nil {
			log.SetLevel(level)
		} else {
			log.Warn("Unrecognized logging verosity value.  Ignoring it.")
		}
	}

	workingDir, _ := os.Getwd()
	app = *config.DefaultApplication(workingDir)
}

func main() {

	if com, ok := operation.GetOperation(&app, operationName); ok {

		com.Init(&app)
		com.Execute(operationFlags...)

	} else {

		log.Error("Unknown operation " + operationName)

	}

}
