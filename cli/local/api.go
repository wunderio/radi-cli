package local

import (
	"os"
	"os/user"
	"path"

	// log "github.com/Sirupsen/logrus"
	// "github.com/urfave/cli"
	"golang.org/x/net/context"

	api_bytesource "github.com/james-nesbitt/wundertools-go/api/handler/bytesource"
	api_local "github.com/james-nesbitt/wundertools-go/api/handler/local"
)

const (
	WUNDERTOOLS_PROJECT_CONF_FOLDER = "wundertools"
	WUNDERTOOLS_USER_CONF_SUBPATH   = "wundertools"
)

func MakeLocalAPI() (*api_local.LocalAPI, error) {

	workingDir, _ := os.Getwd()
	settings := api_local.LocalAPISettings{
		BytesourceFileSettings: api_bytesource.BytesourceFileSettings{
			ExecPath:    workingDir,
			ConfigPaths: &api_bytesource.Paths{},
		},
		Context: context.Background(),
	}

	// Discover paths for the user like ~ and ~/.config/wundertools
	DiscoverUserPaths(&settings)
	DiscoverProjectPaths(&settings)

	/**
	 * We could here add more paths for settings.ConfigPaths, for
	 * configurations of a higher priority.  For example, a feature
	 * or environment concept might want to override user and
	 * project level confs
	 */

	API := api_local.MakeLocalAPI(settings)

	return API, nil
}

// a quick snippet to discover a user's home folder
func userHomePath() string {
	if currentUser, err := user.Current(); err == nil {
		return currentUser.HomeDir
	} else {
		return os.Getenv("HOME")
	}
}

/**
 * Discover some user paths
 *
 * @NOTE we have to play some games for different OSes here
 *
 * dependening on OS, determine if the user has any settings
 * if so, add a conf path for them.
 */
func DiscoverUserPaths(settings *api_local.LocalAPISettings) {
	homeDir := userHomePath()
	homeWTDir := path.Join(homeDir, WUNDERTOOLS_PROJECT_CONF_FOLDER)

	if _, err := os.Stat(path.Join(homeDir, "Library")); err == nil {
		// OSX
		homeWTDir = path.Join(homeDir, "Library", WUNDERTOOLS_USER_CONF_SUBPATH)
	} else if _, err := os.Stat(path.Join(homeDir, ".config")); err == nil {
		// Good Linux
		homeWTDir = path.Join(homeDir, ".config", WUNDERTOOLS_USER_CONF_SUBPATH)
	}

	/**
	 * Set up some frequesntly used paths
	 */
	settings.UserHomePath = homeDir
	settings.ConfigPaths.Set("user-wundertools", homeWTDir)
}

/**
 * Discover project paths
 *
 * Recursively navigate up the file path until we discover a folder that
 * has the key configuration subfolder in it.  That path is marked as the
 * application root, and the subfolder is marked as a conf path
 */
func DiscoverProjectPaths(settings *api_local.LocalAPISettings) {
	workingDir := settings.ExecPath
	homeDir := userHomePath()

	projectRootDirectory := workingDir
	_, err := os.Stat(path.Join(projectRootDirectory, WUNDERTOOLS_PROJECT_CONF_FOLDER))
RootSearch:
	for err != nil {
		projectRootDirectory = path.Dir(projectRootDirectory)
		if projectRootDirectory == homeDir || projectRootDirectory == "." || projectRootDirectory == "/" {
			// Could not find a project folder
			projectRootDirectory = workingDir
			break RootSearch
		}
		_, err = os.Stat(path.Join(projectRootDirectory, WUNDERTOOLS_PROJECT_CONF_FOLDER))
	}

	/**
	 * Set up some frequesntly used paths
	 */
	settings.ProjectRootPath = projectRootDirectory
	settings.ConfigPaths.Set("project-wundertools", path.Join(projectRootDirectory, WUNDERTOOLS_PROJECT_CONF_FOLDER))
}
