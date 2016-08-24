package compose

import (
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

func (project *ComposeProject) Start() {
	if err := project.APIProject.Start(context.Background()); err != nil {
		log.WithError(err).Fatal("Could not start the project.")
	}
}
