package compose

import (
	"golang.org/x/net/context"

	"github.com/docker/libcompose/project/options"
)

func (composeProject *ComposeProject) Down() {
	optionsDown := options.Down{}	

	if err := composeProject.APIProject.Down(context.Background(), optionsDown); err!= nil {
		composeProject.log.Fatal(err.Error())
	}
}
