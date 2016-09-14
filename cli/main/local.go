package main

import (
	"os"
	"os/user"
	"path"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	api_bytesource "github.com/james-nesbitt/wundertools-go/api/handler/bytesource"
	api_local "github.com/james-nesbitt/wundertools-go/api/handler/local"
	api_config "github.com/james-nesbitt/wundertools-go/api/operation/config"
)

const (
	WUNDERTOOLS_PROJECT_CONF_FOLDER = "wundertools"
	WUNDERTOOLS_USER_CONF_SUBPATH   = "wundertools"
)

func TestLocalAPI(c *cli.Context) error {

	workingDir, _ := os.Getwd()
	settings := api_local.LocalAPISettings{ExecPath: workingDir, ConfigPaths: &api_bytesource.Paths{}}

	// Discover paths for the user like ~ and ~/.config/wundertools
	DiscoverUserPaths(&settings)
	DiscoverProjectPaths(&settings)

	/**
	 * We could here add more paths for settings.ConfigPaths, for
	 * configurations of a higher priority.  For example, a feature
	 * or environment concept might want to override user and
	 * project level confs
	 */

	local := api_local.MakeLocalAPI(settings)

	log.WithFields(log.Fields{"api": local}).Info("API")

	log.WithFields(log.Fields{"path": settings.ExecPath}).Info("Exec Path")
	log.WithFields(log.Fields{"path": settings.ProjectRootPath}).Info("Project Root Path")
	log.WithFields(log.Fields{"path": settings.UserHomePath}).Info("User Home Path")

	for index, id := range settings.ConfigPaths.Order() {
		confPath, _ := settings.ConfigPaths.Get(id)
		log.WithFields(log.Fields{"index": index, "id": id, "path": confPath.PathString()}).Info("Config Path: ")
	}

	// get all of the operations
	ops := local.Operations()

	log.Info("Listing operations")
	for _, id := range ops.Order() {
		op, _ := ops.Get(id)

		log.WithFields(log.Fields{"id": op.Id(), "op": op}).Info("Operation: " + op.Label())
		// we could also add "label": op.Label(), "description": op.Description(), "configurations": op.Configurations()
	}

	configList, _ := ops.Get(api_config.OPERATION_ID_CONFIG_LIST)
	KeysConf, _ := configList.Configurations().Get(api_config.OPERATION_CONFIGURATION_CONFIG_KEYS)

	if ok, errs := configList.Exec().Success(); ok {
		keys := KeysConf.Get().([]string)
		log.WithFields(log.Fields{"keys": keys}).Info("Config Key list")

		configGet, _ := ops.Get(api_config.OPERATION_ID_CONFIG_GET)
		valueConf, _ := configGet.Configurations().Get(api_config.OPERATION_CONFIGURATION_CONFIG_VALUE)
		keyConf, _ := configGet.Configurations().Get(api_config.OPERATION_CONFIGURATION_CONFIG_KEY)

		for _, key := range keys {
			keyConf.Set(key)
			configGet.Exec()
			log.WithFields(log.Fields{"key": key, "value": valueConf.Get()}).Info("Config setting: " + key)
		}

	} else {
		logger := log.WithFields(log.Fields{})
		for _, err := range errs {
			logger.WithError(err)
		}
		logger.Error("failed to exec config.get operation")
	}

	configSet, _ := ops.Get(api_config.OPERATION_ID_CONFIG_SET)
	valueConf, _ := configSet.Configurations().Get(api_config.OPERATION_CONFIGURATION_CONFIG_VALUE)
	keyConf, _ := configSet.Configurations().Get(api_config.OPERATION_CONFIGURATION_CONFIG_KEY)

	log.Info("testing config set")
	keyConf.Set("time")
	newValue := time.Now().Format(time.UnixDate)
	valueConf.Set(newValue)
	if ok, errs := configSet.Exec().Success(); ok {
		log.WithFields(log.Fields{"key": "time", "value": newValue}).Info("Saved new key value")
	} else {
		for _, err := range errs {
			log.WithError(err).Error("failed to set config value")
		}
	}

	return nil
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
