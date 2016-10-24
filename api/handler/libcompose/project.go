package libcompose

import (
	"io"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"

	libCompose_docker "github.com/docker/libcompose/docker"
	libCompose_dockerctx "github.com/docker/libcompose/docker/ctx"
	libCompose_project "github.com/docker/libcompose/project"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * A wrapper for the various libCompose objects that
 * are needed to handler orchestration through libcompose
 */

func MakeComposeProject(properties *operation.Properties) (*ComposeProject, bool) {

	projectNameProp, _ := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_PROJECTNAME)
	composeProjectName := projectNameProp.Get().(string)

	projectFilesProp, _ := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_COMPOSEFILES)
	composeFiles := projectFilesProp.Get().([]string)

	contextProp, _ := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_CONTEXT)
	netContext := contextProp.Get().(context.Context)

	outputProp, _ := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_OUTPUT)
	outputWriter := outputProp.Get().(io.Writer)

	errProp, _ := properties.Get(OPERATION_PROPERTY_LIBCOMPOSE_ERROR)
	errorWriter := errProp.Get().(io.Writer)

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

/**
 * clean up a service based on this app
 */
func (project *ComposeProject) AlterService(service *libCompose_config.ServiceConfig) {

	project.alterService_RewriteMappedVolumes(service)

}

// rewrite mapped service volumes to use app points.
func (project *ComposeProject) alterService_RewriteMappedVolumes(service *libCompose_config.ServiceConfig) {

	// short cut to the application paths, which we will use for substitution
	appPaths := project.Paths

	for index, _ := range service.Volumes.Volumes {

		volume := service.Volumes.Volumes[index]

		switch volume.Source[0] {
		/**
		 * @TODO refactor this string comparison to be less cumbersome
		 */
		case []byte("~")[0]:
			homePath, _ := appPaths.Path("user-home")
			volume.Source = strings.Replace(volume.Source, "~", homePath, 1)

		case []byte(".")[0]:
			appPath, _ := appPaths.Path("project-root")
			volume.Source = strings.Replace(volume.Source, "~", appPath, 1)

		case []byte("@")[0]:
			if aliasPath, found := appPaths.Path(volume.Source[1:]); found {
				volume.Source = strings.Replace(volume.Source, volume.Source, aliasPath, 1)
			}
		}
	}

}