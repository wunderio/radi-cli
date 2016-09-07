package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"	

	"github.com/james-nesbitt/wundertools-go/api/handler/null"
)

func TestNullAPI(c *cli.Context) error {

	nAPI := null.MakeNullAPI()

	log.WithFields(log.Fields{"api": nAPI}).Info("API test")

	ops := nAPI.Operations()

	for _, id := range ops.OperationOrder() { 
		op, _ := ops.Operation(id)

		log.WithFields(log.Fields{"id": op.Id()}).Info("Operation: "+op.Label())
		// we could also add "label": op.Label(), "description": op.Description(), "configurations": op.Configurations()
	}

	return nil
}
