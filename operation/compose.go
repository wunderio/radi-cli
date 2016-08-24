package operation

import (
	log "github.com/Sirupsen/logrus"

	"github.com/james-nesbitt/wundertools-go/compose"
)

type Compose struct {
	BaseOperation
}

func (operation *Compose) Execute(flags ...string) {

	composeProject, ok := compose.MakeComposeProject(operation.application)
	if !ok {
		log.Error("could not build compose project")
		return
	}

	if len(flags) > 0 {
		switch flags[0] {
		case "pull":
			log.Debug("Pulling project")
			composeProject.Pull()
		case "up":
			log.Debug("Upping project")
			operation.execute_Up(composeProject, flags...)
		case "down":
			log.Debug("Upping project")
			operation.execute_Down(composeProject, flags...)
		case "start":
			log.Debug("Upping project")
			operation.execute_Start(composeProject, flags...)
		case "stop":
			log.Debug("Upping project")
			operation.execute_Stop(composeProject, flags...)

		case "info":
			log.Debug("Project information")
			composeProject.Info()
		}

	} else {
		log.Warn("No operation was passed to the compose operation")
	}

}

// Parse flags and interpret a compose up
func (operation *Compose) execute_Up(composeProject *compose.ComposeProject, flags ...string) {
	NoRecreate := false
	ForceRecreate := false
	NoBuild := false

	for _, flag := range flags {
		switch flag {
		case "--NoRecreate":
			NoRecreate = true
		case "--Recreate":
			NoRecreate = false
		case "--NoBuild":
			NoBuild = true
		case "--Build":
			NoBuild = false
		case "--ForceRecreate":
			ForceRecreate = true
		}
	}

	composeProject.Up(NoRecreate, ForceRecreate, NoBuild)
}

// Parse flags and interpret a compose down
func (operation *Compose) execute_Down(composeProject *compose.ComposeProject, flags ...string) {
	RemoveVolume := false
	RemoveImages := ""
	RemoveOrphans := true

	for _, flag := range flags {
		switch flag {
		case "--NoRemoveOrphans":
			RemoveOrphans = false
		case "--RemoveVolume":
			RemoveVolume = true
		case "--NoRemoveVolume":
			RemoveVolume = false
		case "--RemoveLocalImages":
			RemoveImages = "local"
		case "--RemoveAllImages":
			RemoveImages = "all"
		}
	}

	composeProject.Down(RemoveVolume, RemoveImages, RemoveOrphans)
}

// Parse flags and interpret a compose start
func (operation *Compose) execute_Start(composeProject *compose.ComposeProject, flags ...string) {
	composeProject.Start()
}

// Parse flags and interpret a compose stop
func (operation *Compose) execute_Stop(composeProject *compose.ComposeProject, flags ...string) {
	timeout := 10

	for _, flag := range flags {
		switch flag {
		case "--quick":
			timeout = 1
		}
	}

	composeProject.Stop(timeout)
}
