package operation

import (
	// "github.com/james-nesbitt/wundertools-go/config"
 	// "github.com/james-nesbitt/wundertools-go/log"
 	"github.com/james-nesbitt/wundertools-go/initialize"
)

type Init struct {
	BaseOperation
}

func (operation *Init) Execute(flags ...string) {

	var method string = "bare"
	var source string = "" // has context based on type

	if len(flags)>0 {
		switch flags[0] {
		case "git":
			method = "git"
			if len(flags)>1 {
				source = flags[0]
			} else {
				operation.logger.Error("No git repository provided.")
			}
		case "yml":
			method = "yml"

			if len(flags)>1 {
				source = flags[0]
			} else {
				operation.logger.Error("No yml source provided.")
			}
		}

	} else {
		operation.logger.Warning("No operation was passed to the compose operation. A bare project will be created.")
	}

	if rootpath, ok := operation.application.Path("project-root"); !ok {
		operation.logger.Error("No project root path has been defined, so no project can be initialized.")
		return
	} else {

		initTasks := initialize.InitTasks{}
		initTasks.Init(operation.logger, rootpath)

		switch method {
		case "git":
			operation.logger.Info("Creating project from git")
			initTasks.Init_Git_Run(operation.logger, source)
		case "bare":
			operation.logger.Info("Creating bare project")
			initTasks.Init_Default_Bare()
		}

		initTasks.RunTasks(operation.logger)

	}

}