package main

import (
	"os"
	"os/user"
	"path"

	handlers_local "github.com/wunderkraut/radi-handlers/local"
)

const (
	RADI_PROJECT_CONF_FOLDER = ".radi" // If the project has existing setitngs, they will be in this subfolder, somewhere up the file tree.
	RADI_USER_CONF_SUBPATH   = "radi"  // If the user has user-scope config, they will be in this subfolder
)

/**
 * Some utility functions for discovering scope paths
 *
 * Scope is usually defined as being in three layers:
 *   - core : stuff that is hardcoded into the API/CLI
 *   - user : configuration for the user across all projects
 *   - project : stuff in the project root configuration path
 *
 *   - environment : specific project environment settings,
 *        usually kept in a sub-path of the project settings.
 *        and keyed using a CLI global flag.
 *
 * @NOTE the difference OSes have different approaches for
 *  determining user settings, usually jsut based on different
 *  paths for the configuration.
 */

// a quick snippet to discover a user's home folder
func userHomePath() string {
	if currentUser, err := user.Current(); err == nil {
		return currentUser.HomeDir
	} else {
		/**
		 * There is an issue in some envs (CoreOSX) where
		 * the golang user library does not set a current user
		 * object properly, so we fall back to checking
		 * an ENV variable.
		 */
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
func DiscoverUserPaths(settings *handlers_local.LocalAPISettings) error {
	var err error

	homeDir := userHomePath()

	// This is a common, but not very good user config path for *Nix OSes
	homeConfDir := path.Join(homeDir, "."+RADI_PROJECT_CONF_FOLDER) // if in the home folder, add a "."

	if _, err = os.Stat(path.Join(homeDir, "Library")); err == nil {
		// OSX
		homeConfDir = path.Join(homeDir, "Library", RADI_USER_CONF_SUBPATH)
	} else if _, err = os.Stat(path.Join(homeDir, ".config")); err == nil {
		// Good *Nix/BSD
		homeConfDir = path.Join(homeDir, ".config", RADI_USER_CONF_SUBPATH)
	}

	/**
	 * @TODO does anybody care about any other OS?
	 */

	/**
	 * Set up some frequesntly used paths
	 */
	settings.UserHomePath = homeDir
	settings.ConfigPaths.Set("user", homeConfDir)

	return err
}

/**
 * Discover project paths
 *
 * Recursively navigate up the file path until we discover a folder that
 * has the key configuration subfolder in it.  That path is marked as the
 * application root, and the subfolder is marked as a conf path
 */
func DiscoverProjectPaths(settings *handlers_local.LocalAPISettings) error {
	workingDir := settings.ExecPath
	homeDir := userHomePath()

	projectRootDirectory := workingDir
	_, err := os.Stat(path.Join(projectRootDirectory, RADI_PROJECT_CONF_FOLDER))
RootSearch:
	for err != nil {
		projectRootDirectory = path.Dir(projectRootDirectory)
		if projectRootDirectory == homeDir || projectRootDirectory == "." || projectRootDirectory == "/" {
			// Could not find a project folder
			projectRootDirectory = workingDir
			settings.ProjectDoesntExist = true
			break RootSearch
		}
		_, err = os.Stat(path.Join(projectRootDirectory, RADI_PROJECT_CONF_FOLDER))
	}

	/**
	 * Set up some frequesntly used paths
	 */
	settings.ProjectRootPath = projectRootDirectory
	settings.ConfigPaths.Set("project", path.Join(projectRootDirectory, RADI_PROJECT_CONF_FOLDER))

	return err
}

/**
 * Discover environment path for a specific environment
 *
 */
func DiscoverEnvironmentPath(settings *handlers_local.LocalAPISettings, environment string) error {

	/**
	 * @TODO actually check to see if the path exists, so that we can warn if it doesn't?
	 */

	// add the environment sub path of the main project conf directory, as a conf path
	settings.ConfigPaths.Set(environment, path.Join(settings.ProjectRootPath, RADI_PROJECT_CONF_FOLDER, environment))

	return nil
}
