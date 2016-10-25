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
			alias := id[strings.Index(id, ".")+1:]
			opWrapper := CliOperationWrapper{op: op}

			cliComm := cli.Command{
				Name:     op.Id(),
				Aliases:  []string{alias},
				Usage:    op.Description(),
				Action:   opWrapper.Exec,
				Category: category,
			}

			cliComm.Flags = CliMakeFlagsFromProperties(*op.Properties())

			app.Commands = append(app.Commands, cliComm)
		}
	}

	if comList, err := local.Command.List(""); err == nil {
		category := "custom"

		for _, key := range comList {
			comm, _ := local.Command.Get(key)

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
	logger := log.WithFields(log.Fields{"id": opWrapper.op.Id()})
	logger.Debug("Running operation")

	props := opWrapper.op.Properties()

	CliAssignPropertiesFromFlags(cliContext, props)

	var success bool
	var errs []error

	if success, errs = opWrapper.op.Exec().Success(); !success {
		if len(errs) == 0 {
			errs = []error{errors.New("Unknown error occured")}
		}
	}

	// Create some meaningful output, by logging some of the properties
	fields := map[string]interface{}{
		"success": success,
		"errors":  errs,
	}
	for _, key := range props.Order() {
		prop, _ := props.Get(key)

		if !prop.Internal() {
			switch prop.Type() {
			case "string":
				fields[key] = prop.Get().(string)
			case "[]string":
				fields[key] = prop.Get().([]string)
			case "[]byte":
				fields[key] = string(prop.Get().([]byte))
			case "int32":
				fields[key] = int(prop.Get().(int32))
			case "int64":
				fields[key] = prop.Get().(int64)
			case "bool":
				fields[key] = prop.Get().(bool)
			}
		}
	}
	logger = logger.WithFields(log.Fields(fields))

	if len(errs) > 0 {
		for _, err := range errs {
			logger = logger.WithError(err)
		}

		logger.Error("Error occured running operation")
		return errs[0]
	} else {
		logger.Info("Operation completed.")
		return nil
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
