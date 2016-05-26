package compose

/**
 * Wrapper for libCompose
 */

import (
	"path"

	"github.com/james-nesbitt/wundertools-go/config"
	"github.com/james-nesbitt/wundertools-go/log"

	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/project"
)

func MakeComposeProject(logger log.Log, application *config.Application) (*ComposeProject, bool) {

	composeProjectName := application.Name
	composeFiles := []string{}


	if yamlPath, ok := application.Paths.Path("project-wundertools"); ok {
		yamlPath = path.Join(yamlPath, "docker-compose.yml")
		composeFiles = append(composeFiles, yamlPath)
	}

	project, err := docker.NewProject(&docker.Context{
		Context: project.Context{
			ComposeFiles: composeFiles,
			ProjectName:  composeProjectName,
		},
	}, nil)

	if err != nil {
		logger.Fatal(err.Error())
		return nil, false
	}

	composeProject := ComposeProject{
		log: logger,
		application: application,
		APIProject: project,
	}

	return &composeProject, true
}

// A wundertools wrapper for the APIProject class
type ComposeProject struct {
	log log.Log
	application *config.Application

	project.APIProject
}
