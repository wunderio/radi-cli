package compose

import (
	"golang.org/x/net/context"

	"github.com/docker/libcompose/project/options"
)

func (composeProject *ComposeProject) Up() {
	optionsUp := options.Up{}	

	if err := composeProject.APIProject.Up(context.Background(), optionsUp); err!= nil {
		composeProject.log.Fatal(err.Error())
	}
}
