package compose

/**
 * Wrapper for libCompose
 */

import (
	"path"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"

	// libCompose_config "github.com/docker/libcompose/config"
	libCompose_docker "github.com/docker/libcompose/docker"
	libCompose_project "github.com/docker/libcompose/project"

	"github.com/james-nesbitt/wundertools-go/config"
)

func MakeComposeProject(application *config.Application) (*ComposeProject, bool) {

	composeProjectName := application.Name
	composeFiles := []string{}

	if yamlPath, ok := application.Paths.Path("project-root"); ok {
		yamlPath = path.Join(yamlPath, "docker-compose.yml")
		composeFiles = append(composeFiles, yamlPath)
	}

	context := &libCompose_docker.Context{
		Context: libCompose_project.Context{
			ComposeFiles: composeFiles,
			ProjectName:  composeProjectName,
		},
	}
	project, err := libCompose_docker.NewProject(context, nil)

	if err != nil {
		log.WithError(err).Fatal("Could not make the docker-compose project.")
		return nil, false
	}

	composeProject := ComposeProject{
		application: application,
		services:    []string{},
		context:     context,
		APIProject:  project,
	}

	return &composeProject, true
}

// A wundertools wrapper for the APIProject class
type ComposeProject struct {
	application *config.Application

	services []string

	context *libCompose_docker.Context
	libCompose_project.APIProject
}

// get a specific service
func (project *ComposeProject) Service(name string) (libCompose_project.Service, error) {
	return project.context.Project.CreateService(name)
}

// List all the service names
func (project *ComposeProject) serviceNames() []string {
	return project.context.Project.ServiceConfigs.Keys()
}

func (project *ComposeProject) ServicePS(names ...string) (libCompose_project.InfoSet, error) {
	return project.context.Project.Ps(context.Background(), false, names...)
}
