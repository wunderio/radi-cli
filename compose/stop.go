package compose

import (
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

func (project *ComposeProject) Stop(timeout int) {
	if err := project.APIProject.Stop(context.Background(), timeout); err != nil {
		log.WithError(err).Fatal("Could not stop the project.")
	}
}
