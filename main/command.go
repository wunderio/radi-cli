package main

import (
	"errors"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"

	api_command "github.com/wunderkraut/radi-api/operation/command"
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
 *  retrieving wrappers from api structs? both approaches
 *  have advantages
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

			cliComm.Flags = CliMakeFlagsFromProperties(comm.Properties())

			log.WithFields(log.Fields{"id": comm.Id()}).Debug("CLI: Adding API command")
			app.Commands = append(app.Commands, &cliComm)

		}

		return err
	} else {
		log.WithError(err).Debug("Failed to list commands")
		return err
	}
}

/**
 * Wrapper for command Exec methods, from the urface CLI
 *
 * The job of this struct is to wrap an API command struct in
 * a format that meets the requirements of the urfave.CLI library
 * mainly by translating the .Exec() method arguments and output.
 */
type CliCommandWrapper struct {
	comm api_command.Command
}

// Execute the operation for the cli
func (commWrapper *CliCommandWrapper) Exec(cliContext *cli.Context) error {
	logger := log.WithFields(log.Fields{"id": commWrapper.comm.Id()})
	logger.Debug("Running command")

	comm := commWrapper.comm
	props := comm.Properties()
	CliAssignPropertiesFromFlags(cliContext, &props)

	// if there was a command flags property, then add any remaining arguments as flags
	if flagsProp, found := props.Get(api_command.OPERATION_PROPERTY_COMMAND_FLAGS); found {
		flagsProp.Set([]string(cliContext.Args().Slice()))
	}

	result := comm.Exec(&props)
	<-result.Finished()

	if !result.Success() {
		var err error
		errs := result.Errors()

		if len(errs) > 0 {
			for _, err := range errs {
				logger = logger.WithError(err)
			}
			err = errs[0]
		} else {
			err = errors.New("Command failed to execute for unknown reasons.")
			logger = logger.WithError(err)
		}

		logger.Error("Error occured running command")
		return err
	}
	return nil
}
