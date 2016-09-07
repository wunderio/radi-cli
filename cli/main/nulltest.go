package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"	

	"github.com/james-nesbitt/wundertools-go/api/handler/null"
	"github.com/james-nesbitt/wundertools-go/api/operation/monitor"
)

func TestNullAPI(c *cli.Context) error {

	nAPI := null.MakeNullAPI()

	log.WithFields(log.Fields{"api": nAPI}).Info("Null API test")

	ops := nAPI.Operations()

	for _, id := range ops.Order() { 
		op, _ := ops.Get(id)

		log.WithFields(log.Fields{"id": op.Id()}).Info("Operation: "+op.Label())
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

	return nil
}
