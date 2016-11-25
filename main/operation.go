package main

import (
	"errors"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli.v2"

	api_operation "github.com/james-nesbitt/kraut-api/operation"
	api_security "github.com/james-nesbitt/kraut-api/operation/security"
)

// Add operations from the API to the app
func AppApiOperations(app *cli.App, ops api_operation.Operations) error {
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

			log.WithFields(log.Fields{"id": id}).Debug("Cli: Adding Operation")
			app.Commands = append(app.Commands, &cliComm)
		}
	}

	return nil
}

/**
 * Wrapper for operation Exec methods, from the urface CLI
 *
 * We use this wrapper because:
 *  1. the cli library has a different return
 *     expectation than what our operations return
 *  2. we need to do some minor transformation on CLI
 *     arguments, to make them fit our types.
 *  3. we want to do some work to decide what to output
 *     to the screen.
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
			errs = []error{errors.New("KrautCLI: Unknown error occured")}
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
			case "github.com/james-nesbitt/kraut-api/operation/security.SecurityUser":
				user := prop.Get().(api_security.SecurityUser)
				fields[key] = user.Id()
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
