package local

import (
	"errors"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	api_operation "github.com/james-nesbitt/wundertools-go/api/operation"
	api_command "github.com/james-nesbitt/wundertools-go/api/operation/command"
)

/**
 * Add Local Commands to the app
 */
func AppLocalCommands(app *cli.App) error {

	local, err := MakeLocalAPI()

	// get all of the operations
	ops := local.Operations()

	for _, id := range ops.Order() {
		op, _ := ops.Get(id)

		log.WithFields(log.Fields{"id": op.Id()}).Debug("Operation: " + op.Label())
		// we could also add "label": op.Label(), "description": op.Description(), "configurations": op.Properties()

		if !op.Internal() {
			id := op.Id()
			category := id[0:strings.Index(id, ".")]
			opWrapper := CliOperationWrapper{op: op}

			cliComm := cli.Command{
				Name:     op.Id(),
				Usage:    op.Description(),
				Action:   opWrapper.Exec,
				Category: category,
			}

			cliComm.Flags = CliMakeFlagsFromProperties(*op.Properties())

			app.Commands = append(app.Commands, cliComm)
		}
	}

	if comList, err := local.Command.List(""); err == nil {
		category := "commands"

		for _, key := range comList {
			comm, _ := local.Command.Get(key)

			commWrapper := CliCommandWrapper{comm: comm}

			cliComm := cli.Command{
				Name:     comm.Id(),
				Usage:    comm.Description(),
				Action:   commWrapper.Exec,
				Category: category,
			}

			cliComm.Flags = CliMakeFlagsFromProperties(*comm.Properties())

			app.Commands = append(app.Commands, cliComm)

		}
	} else {
		log.WithError(err).Error("Failed to list commands")
	}

	return err
}

/**
 * Wrapper for operation Exec methods, from the urface CLI
 */
type CliOperationWrapper struct {
	op api_operation.Operation
}

// Execute the operation for the cli
func (opWrapper *CliOperationWrapper) Exec(cliContext *cli.Context) error {
	log.WithFields(log.Fields{"id": opWrapper.op.Id()}).Debug("Running operation")

	CliAssignPropertiesFromFlags(cliContext, opWrapper.op.Properties())

	if success, errs := opWrapper.op.Exec().Success(); !success {
		var err error
		if len(errs) > 0 {
			err = errs[0]
		} else {
			err = errors.New("Unknown error occured")
		}
		log.WithError(err).Error("Error occured running operation")
	}
	return nil
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
