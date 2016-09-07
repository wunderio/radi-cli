package main

import (
	"os"

	"github.com/urfave/cli"
	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/version"
	"github.com/james-nesbitt/wundertools-go/api/handler/null"
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
			Name:    "test",
			Usage:   "Test CLI",
			Action:  TestCommand,
		},
		{
			Name:    "api",
			Usage:   "Test API",
			Action:  TestAPI,
		},
	}

	app.Run(os.Args)

}

func TestCommand(c *cli.Context) error {
	log.WithFields(log.Fields{"args": c.Args()}).Info("added task")
	return nil
}

func TestAPI(c *cli.Context) error {

	nAPI := null.MakeNullAPI()

	log.WithFields(log.Fields{"api": nAPI}).Info("API test")


	return nil
}