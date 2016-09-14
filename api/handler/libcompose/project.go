package libcompose

import (
	"io"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"

	// libCompose_config "github.com/docker/libcompose/config"
	libCompose_docker "github.com/docker/libcompose/docker"
	libCompose_dockerctx "github.com/docker/libcompose/docker/ctx"
	libCompose_project "github.com/docker/libcompose/project"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * A wrapper for the various libCompose objects that
 * are needed to handler orchestration through libcompose
 */

func MakeComposeProject(configurations *operation.Configurations) (*ComposeProject, bool) {

	projectNameConf, _ := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_PROJECTNAME)
	composeProjectName := projectNameConf.Get().(string)

	projectFilesConf, _ := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_COMPOSEFILES)
	composeFiles := projectFilesConf.Get().([]string)

	contextConf, _ := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_CONTEXT)
	netContext := contextConf.Get().(context.Context)

	outputConf, _ := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_OUTPUT)
	outputWriter := outputConf.Get().(io.Writer)

	errConf, _ := configurations.Get(OPERATION_CONFIGURATION_LIBCOMPOSE_ERROR)
	errorWriter := errConf.Get().(io.Writer)

	loggerFactory := NewLibcomposeLoggerFactory(outputWriter, errorWriter)

	composeContext := &libCompose_dockerctx.Context{
		Context: libCompose_project.Context{
			ComposeFiles:  composeFiles,
			ProjectName:   composeProjectName,
			LoggerFactory: loggerFactory,
		},
	}

	project, err := libCompose_docker.NewProject(composeContext, nil)

	if err != nil {
		log.WithError(err).Fatal("Could not make the docker-compose project.")
		return nil, false
	}

	composeProject := ComposeProject{
		netContext:     netContext,
		composeContext: composeContext,
		APIProject:     project,
	}

	return &composeProject, true
}

// A wundertools wrapper for the APIProject class
type ComposeProject struct {
	libCompose_project.APIProject
	netContext     context.Context
	composeContext *libCompose_dockerctx.Context
}

// get a specific service
func (project *ComposeProject) Service(name string) (libCompose_project.Service, error) {
	return project.composeContext.Project.CreateService(name)
}

// List all the service names
func (project *ComposeProject) serviceNames() []string {
	return project.composeContext.Project.ServiceConfigs.Keys()
}

func (project *ComposeProject) ServicePS(names ...string) (libCompose_project.InfoSet, error) {
	return project.composeContext.Project.Ps(context.Background(), names...)
}

func (project *ComposeProject) Context() context.Context {
	return project.netContext
}
