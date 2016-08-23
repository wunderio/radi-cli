package operation

import (
	"io"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"

	// "github.com/james-nesbitt/wundertools-go/config"
	"github.com/james-nesbitt/wundertools-go/initialize"
)

type InitGenerate struct {
	BaseOperation
}

func (operation *InitGenerate) Execute(flags ...string) {

	var method string = "yaml"
	var writer io.Writer

	skip := []string{}

	if method == "test" {
		logger := log.StandardLogger().Writer()
		defer logger.Close()
		writer = io.Writer(logger)
	} else {
		destination, _ := operation.application.Path("project-wundertools")
		destination = path.Join(destination, "init.yml")

		if file, err := os.Create(destination); err == nil {
			skip = append(skip, ".wundertools/init.yml")
			defer file.Close()
			writer = io.Writer(file)
		}
	}

	if rootpath, ok := operation.application.Path("project-root"); !ok {
		log.Error("No project root path has been defined, so no project can be initialized.")
		return
	} else {

		initialize.Init_Generate(method, rootpath, skip, 1024*1024, writer)

	}

}
