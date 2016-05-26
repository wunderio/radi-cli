package command

import (
	"github.com/james-nesbitt/wundertools-go/compose"
// 	"github.com/james-nesbitt/wundertools-go/config"
// 	"github.com/james-nesbitt/wundertools-go/log"
)

type Compose struct {
	BaseCommand
}

func (command *Compose) Execute(flags []string) {

	composeProject, ok := compose.MakeComposeProject(command.logger, command.application)
	if !ok {
		command.logger.Error("could not build compose project")
		return
	}

	if len(flags)>0 {
		switch flags[0] {
		case "up":
			composeProject.Up()
		}

	}

}
