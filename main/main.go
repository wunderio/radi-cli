package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

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

	/**
	 * As a shortcut, we are wiring the CLI to use just a
	 * local API implementation.  This hand-off function does
	 * that.
	 *
	 * @TODO write a better implementation for this
	 */
	AppLocalCommands(app)

	app.Run(os.Args)

}
