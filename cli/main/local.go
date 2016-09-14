package main

import (
	"net/http"
	"os"
	"os/user"
	"path"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"golang.org/x/net/context"

	api_bytesource "github.com/james-nesbitt/wundertools-go/api/handler/bytesource"
	api_libcompose "github.com/james-nesbitt/wundertools-go/api/handler/libcompose"
	api_local "github.com/james-nesbitt/wundertools-go/api/handler/local"
	api_config "github.com/james-nesbitt/wundertools-go/api/operation/config"
	api_orchestrate "github.com/james-nesbitt/wundertools-go/api/operation/orchestrate"
)

const (
	WUNDERTOOLS_PROJECT_CONF_FOLDER = "wundertools"
	WUNDERTOOLS_USER_CONF_SUBPATH   = "wundertools"
)

func TestLocalAPI(c *cli.Context) error {

	workingDir, _ := os.Getwd()
	settings := api_local.LocalAPISettings{
		ExecPath:    workingDir,
		ConfigPaths: &api_bytesource.Paths{},
		Context:     context.Background(),
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

		log.WithFields(log.Fields{"id": op.Id()}).Info("Operation: " + op.Label())
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
		logger.Error("failed to exec config.list operation")
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

		configGet, _ := ops.Get(api_config.OPERATION_ID_CONFIG_GET)
		valueConf, _ := configGet.Configurations().Get(api_config.OPERATION_CONFIGURATION_CONFIG_VALUE)
		keyConf, _ := configGet.Configurations().Get(api_config.OPERATION_CONFIGURATION_CONFIG_KEY)

		key := keyConf.Get().(string)
		keyConf.Set(key)
		configGet.Exec()
		log.WithFields(log.Fields{"key": key, "value": valueConf.Get()}).Info("Checking new key value")
	} else {
		for _, err := range errs {
			log.WithError(err).Error("failed to set config value")
		}
	}

	// Before testing orchestration, let's attach to log output
	log.Info("attaching to log output before testing orchestration")

	monitorLogs, _ := ops.Get(api_libcompose.OPERATION_ID_COMPOSE_MONITOR_LOGS)

	outputConf, _ := monitorLogs.Configurations().Get(api_libcompose.OPERATION_CONFIGURATION_LIBCOMPOSE_OUTPUT)
	outputConf.Set(os.Stdout)

	stayAttachedConf, _ := monitorLogs.Configurations().Get(api_libcompose.OPERATION_CONFIGURATION_LIBCOMPOSE_ATTACH_FOLLOW)
	if !stayAttachedConf.Set(false) {
		log.Error("Could not set logs to follow")
	}

	// This runs logging as an inline action, without following, which will output any existing log items
	// this will output somethign if the containers already exist.
	log.Info("output logs before starting orchestration tests")
	if success, errs := monitorLogs.Exec().Success(); !success {
		for _, err := range errs {
			log.WithError(err).Error("Logs operation failed")
		}
	}
	log.Info("logs outputted")

	// test orchestrate up
	log.Info("Testing orchestration up")
	orchestrateUp, _ := ops.Get(api_orchestrate.OPERATION_ID_ORCHESTRATE_UP)
	if success, errs := orchestrateUp.Exec().Success(); !success {
		for _, err := range errs {
			log.WithError(err).Error("failed to run orchestrate up")
		}
	}

	// run the log tracking execution in it's own thread
	// it appears that this stay open until the containers are shut down
	if !stayAttachedConf.Set(true) {
		log.Error("Could not set logs to follow")
	}
	go func() {
		os.Stdout.Write([]byte("Log follow started\n"))
		if success, _ := monitorLogs.Exec().Success(); !success {
			log.Error("Couldn't run log monitor")
		}
		os.Stdout.Write([]byte("Log follow ended\n"))
	}()

	response, err := http.Get("http://wundertools.docker")
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
	}

	// test orchestrate up
	log.Info("Testing orchestration down")
	orchestrateDown, _ := ops.Get(api_orchestrate.OPERATION_ID_ORCHESTRATE_DOWN)
	if success, errs := orchestrateDown.Exec().Success(); !success {
		for _, err := range errs {
			log.WithError(err).Error("failed to run orchestrate down")
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
