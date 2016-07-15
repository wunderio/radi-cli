package compose

import (
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

func (project *ComposeProject) Pull(services ...string) {
	if err := project.APIProject.Pull(context.Background()); err != nil {
		log.WithError(err).Panic("Could not pull the project.")
	}
}
