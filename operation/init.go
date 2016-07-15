package operation

import (
	log "github.com/Sirupsen/logrus"

	// "github.com/james-nesbitt/wundertools-go/config"
	"github.com/james-nesbitt/wundertools-go/initialize"
)

type Init struct {
	BaseOperation
}

func (operation *Init) Execute(flags ...string) {

	var method string = "bare"
	var source string = "" // has context based on type

	if len(flags) > 0 {
		switch flags[0] {
		case "git":
			method = "git"
			if len(flags) > 1 {
				source = flags[0]
			} else {
				log.Error("No git repository provided.")
			}
		case "yml":
			method = "yml"

			if len(flags) > 1 {
				source = flags[0]
			} else {
				log.Error("No yml source provided.")
			}
		}

	} else {
		log.Warning("No operation was passed to the compose operation. A bare project will be created.")
	}

	if rootpath, ok := operation.application.Path("project-root"); !ok {
		log.Error("No project root path has been defined, so no project can be initialized.")
		return
	} else {

		initTasks := initialize.InitTasks{}
		initTasks.Init(rootpath)

		switch method {
		case "git":
			log.Info("Creating project from git")
			initTasks.Init_Git_Run(source)
		case "bare":
			log.Info("Creating bare project")
			initTasks.Init_Default_Bare()
		}

		initTasks.RunTasks()

	}

}
