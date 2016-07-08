package compose

import (
	"golang.org/x/net/context"
)

func (project *ComposeProject) Pull(services ...string) {
	if err := project.APIProject.Pull(context.Background()); err != nil {
		project.log.Fatal(err.Error())
	}
}
