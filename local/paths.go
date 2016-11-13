package local

import (
	"os"
	"os/user"
	"path"

	handlers_local "github.com/james-nesbitt/kraut-handlers/local"
)

/**
 * Some utility functions for discovering paths related
 * to local proect settings
 */

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
func DiscoverUserPaths(settings *handlers_local.LocalAPISettings) error {
	homeDir := userHomePath()

	// This is a common, but not very good user config path for *Nix OSes
	homeConfDir := path.Join(homeDir, "."+WUNDERTOOLS_PROJECT_CONF_FOLDER) // if in the home folder, add a "."

	if _, err := os.Stat(path.Join(homeDir, "Library")); err == nil {
		// OSX
		homeConfDir = path.Join(homeDir, "Library", WUNDERTOOLS_USER_CONF_SUBPATH)
	} else if _, err := os.Stat(path.Join(homeDir, ".config")); err == nil {
		// Good *Nix/BSD
		homeConfDir = path.Join(homeDir, ".config", WUNDERTOOLS_USER_CONF_SUBPATH)
	}

	/**
	 * @TODO does anybody care about any other OS?
	 */

	/**
	 * Set up some frequesntly used paths
	 */
	settings.UserHomePath = homeDir
	settings.ConfigPaths.Set("user", homeConfDir)

	return nil
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
	_, err := os.Stat(path.Join(projectRootDirectory, WUNDERTOOLS_PROJECT_CONF_FOLDER))
RootSearch:
	for err != nil {
		projectRootDirectory = path.Dir(projectRootDirectory)
		if projectRootDirectory == homeDir || projectRootDirectory == "." || projectRootDirectory == "/" {
			// Could not find a project folder
			projectRootDirectory = workingDir
			settings.ProjectDoesntExist = true
			break RootSearch
		}
		_, err = os.Stat(path.Join(projectRootDirectory, WUNDERTOOLS_PROJECT_CONF_FOLDER))
	}

	/**
	 * Set up some frequesntly used paths
	 */
	settings.ProjectRootPath = projectRootDirectory
	settings.ConfigPaths.Set("project", path.Join(projectRootDirectory, WUNDERTOOLS_PROJECT_CONF_FOLDER))

	return nil
}
