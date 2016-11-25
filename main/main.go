package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli.v2"

	api_command "github.com/james-nesbitt/kraut-api/operation/command"
	cli_local "github.com/james-nesbitt/kraut-cli/local"
	"github.com/james-nesbitt/kraut-cli/version"
)

func main() {
	var debug bool

	app := &cli.App{}

	app.Name = "wundertools"
	app.Usage = "Command line interface for Kraut API."
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"
	app.Authors = []*cli.Author{&cli.Author{Name: "Wunder.IO", Email: "https://github.com/james-nesbitt/kraut-cli"}}

	app.Flags = []cli.Flag{
		cli.Flag(cli.Flag(&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Enable verbose debugging output",
			EnvVars:     []string{"KRAUT_DEBUG"},
			Hidden:      false,
			Destination: &debug,
		})),
	}

	// Run these functions before processing
	app.Before = func(c *cli.Context) error {
		if err := Before_GlobalFlags(c); err != nil {
			return err
		}
		return nil
	}

	// Make a local API instance
	local, _ := cli_local.MakeLocalAPI()

	// Catch the ops
	localOps := local.Operations()

	// Add any operations from the api to the app
	AppApiOperations(app, localOps)

	// Add any commands from the api CommandWrapper to the app
	AppWrapperCommands(app, api_command.New_SimpleCommandWrapper(&localOps))

	// Run the App initializer
	app.Setup()

	if debug {

		/**
		 * @TODO do something here to make logrus output debug
		 *  statements
		 */
	}

	// Run the CLI command based on passed args
	app.Run(os.Args)
}

/**
 * Before functions
 */

// Process global flags
func Before_GlobalFlags(c *cli.Context) error {
	if c.IsSet("debug") {
		log.SetLevel(log.DebugLevel)
		log.Debug("Enabling verbose debug output")
	}
	return nil
}
