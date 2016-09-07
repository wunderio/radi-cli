package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/james-nesbitt/wundertools-go/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "wundertools"
	app.Usage = "Command line interface for Wundertools API."
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"
	app.Author = "Wunder.IO"
	app.Email = "https://github.com/james-nesbitt/wundertools-go"
	//	app.Flags = append(command.CommonFlags(), dockerApp.DockerClientFlags()...)
	app.Commands = []cli.Command{
		{
			Name:   "test",
			Usage:  "Test CLI",
			Action: TestCommand,
		},
		{
			Name:   "null-api",
			Usage:  "Test API",
			Action: TestNullAPI,
		},
	}

	app.Run(os.Args)

}

func TestCommand(c *cli.Context) error {
	log.WithFields(log.Fields{"args": c.Args()}).Info("added task")
	return nil
}
