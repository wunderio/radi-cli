package clitest

import (
	// log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

/**
 * Add Testing Commands to the app
 */
func AppAddTests(app *cli.App) {

	app.Commands = append(app.Commands, cli.Command{
		Name:   "null-api",
		Usage:  "Test Null API",
		Action: TestNullAPI,
	})
	app.Commands = append(app.Commands, cli.Command{
		Name:   "local-api",
		Usage:  "Test Local API",
		Action: TestLocalAPI,
	})

}
