package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/james-nesbitt/kraut-cli/version"

	cli_local "github.com/james-nesbitt/kraut-cli/local"
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
	}

	// Add local commands
	cli_local.AppLocalCommands(app)

	app.Run(os.Args)

}
