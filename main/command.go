package main

import (
	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	api_command "github.com/james-nesbitt/kraut-api/operation/command"
)

/**
 * Add any commands from a command wrapper to the app
 *
 * @note you can get a CommandWrapper from certain APIs, or
 *   you can build one yourself by passing in any list of
 *   operations into one of the CommandWrapper constructors.
 *
 * @todo should this receive an API list of operations and
 *  build it's own command wrapper, or should we abstract
 *  retrieving wrappers from api structs?
 */
func AppWrapperCommands(app *cli.App, commands api_command.CommandWrapper) error {
	if comList, err := commands.List(""); err == nil {
		category := "custom"

		for _, key := range comList {
			comm, _ := commands.Get(key)

			commWrapper := CliCommandWrapper{comm: comm}

			cliComm := cli.Command{
				Name:     "command.exec." + comm.Id(),
				Aliases:  []string{comm.Id()},
				Usage:    comm.Description(),
				Action:   commWrapper.Exec,
				Category: category,
			}

			cliComm.Flags = CliMakeFlagsFromProperties(*comm.Properties())

			app.Commands = append(app.Commands, cliComm)

		}

		return err
	} else {
		log.WithError(err).Error("Failed to list commands")
		return err
	}
}

/**
 * Wrapper for command Exec methods, from the urface CLI
 */
type CliCommandWrapper struct {
	comm api_command.Command
}

// Execute the operation for the cli
func (commWrapper *CliCommandWrapper) Exec(cliContext *cli.Context) error {
	log.WithFields(log.Fields{"id": commWrapper.comm.Id()}).Debug("Running command")

	CliAssignPropertiesFromFlags(cliContext, commWrapper.comm.Properties())

	// if there was a command flags property, then add any remaining arguments as flags
	if flagsProp, found := commWrapper.comm.Properties().Get(api_command.OPERATION_PROPERTY_COMMAND_FLAGS); found {
		flagsProp.Set([]string(cliContext.Args()))
	}

	if success, errs := commWrapper.comm.Exec().Success(); !success {
		var err error
		if len(errs) > 0 {
			err = errs[0]
		} else {
			err = errors.New("Unknown error occured")
		}
		log.WithError(err).Error("Error occured running command")
	}
	return nil
}
