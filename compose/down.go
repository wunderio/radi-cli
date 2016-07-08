package compose

import (
	"golang.org/x/net/context"

	"github.com/docker/libcompose/project/options"
)

func (project *ComposeProject) Down(RemoveVolume bool, RemoveImages string, RemoveOrphans bool) {
	optionsDown := options.Down{
		RemoveVolume:  RemoveVolume,
		RemoveImages:  options.ImageType(RemoveImages),
		RemoveOrphans: RemoveOrphans,
	}

	if err := project.APIProject.Down(context.Background(), optionsDown); err != nil {
		project.log.Fatal(err.Error())
	}
}
