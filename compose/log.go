package compose

import (
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

func (project *ComposeProject) Log(follow bool) {
	if err := project.APIProject.Log(context.TODO(), follow); err != nil {
		log.WithError(err).Fatal("Could not start the project.")
	}
}
