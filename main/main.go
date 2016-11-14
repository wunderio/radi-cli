package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	api_command "github.com/james-nesbitt/kraut-api/operation/command"
	cli_local "github.com/james-nesbitt/kraut-cli/local"
	"github.com/james-nesbitt/kraut-cli/version"
)

func main() {
	var debug bool

	app := cli.NewApp()
	app.Name = "wundertools"
	app.Usage = "Command line interface for Kraut API."
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"
	app.Author = "Wunder.IO"
	app.Email = "https://github.com/james-nesbitt/kraut-cli"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable verbose debugging output",
			EnvVar:      "KRAUT_DEBUG",
			Hidden:      false,
			Destination: &debug,
		},
	}

	if debug {
		log.Info("Enabling verbose debug output")

		/**
		 * @TODO do something here to make logrus output debug
		 *  statements
		 */
	}

	// Make a local API instance
	local, _ := cli_local.MakeLocalAPI()

	// Catch the ops
	localOps := local.Operations()

	// Add any operations from the api to the app
	AppApiOperations(app, localOps)

	// Add any commands from the api CommandWrapper to the app
	AppWrapperCommands(app, api_command.New_SimpleCommandWrapper(&localOps))

	// Run the CLI command based on passed args
	app.Run(os.Args)
}
