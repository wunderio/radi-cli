package main

import (
	// "text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/james-nesbitt/wundertools-go/api/handler/null"
	"github.com/james-nesbitt/wundertools-go/api/operation/monitor"
)

func TestNullAPI(c *cli.Context) error {

	// occasionally, throughout this function, we are going to
	// want to just write something to output, but we will still
	// pipe it out through the logger
	writer := log.StandardLogger().Writer()
	defer writer.Close()

	nAPI := null.MakeNullAPI()

	log.WithFields(log.Fields{"api": nAPI}).Info("Null API test")

	// get all of the operations
	ops := nAPI.Operations()

	// let's get the logging operation, which we may want to use repeatedly
	logger, _ := ops.Get(monitor.OPERATION_ID_MONITOR_LOG)

	// just for fun, let's output some information about the operation
	log.WithFields(log.Fields{"id": logger.Id(), "label": logger.Label()}).Info("Logger operation found")
	loggerConfs := logger.Configurations()
	for _, id := range loggerConfs.Order() {
		conf, _ := loggerConfs.Get(id)
		log.WithFields(log.Fields{"conf": conf}).Info("Logger has conf field: " + id)
	}

	log.Info("Listing operations")
	for _, id := range ops.Order() {
		op, _ := ops.Get(id)

		log.WithFields(log.Fields{"id": op.Id()}).Info("Operation: " + op.Label())
		// we could also add "label": op.Label(), "description": op.Description(), "configurations": op.Configurations()
	}

	// if there is a monitor.status operation, attach a writer to it and use it
	if status, exists := ops.Get(monitor.OPERATION_ID_MONITOR_STATUS); exists {
		log.Info("Running status operation")
		status.Exec()
	} else {
		log.Warning("No status operations was available")
	}

	// if there is a monitor.info operation, attach a writer to it and use it
	if info, exists := ops.Get(monitor.OPERATION_ID_MONITOR_INFO); exists {
		log.Info("Running info operation")
		info.Exec()
	} else {
		log.Warning("No info operations was available")
	}

	// lets try directly using the log operation (just ignoring the Get() bool for now)
	confs := logger.Configurations()
	logtype, _ := confs.Get(monitor.OPERATION_CONFIGURATION_CONF_MONITOR_LOG_TYPE)
	message, _ := confs.Get(monitor.OPERATION_CONFIGURATION_CONF_MONITOR_LOG_MESSAGE)

	// start off with just info
	logtype.Set("info") // this is actually default
	message.Set("here is my test info")
	logger.Exec()
	message.Set("here is my second test info")
	logger.Exec()
	// Now try an error
	logtype.Set("error")
	message.Set("here is my test Error")
	logger.Exec()

	return nil
}
