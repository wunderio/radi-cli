package config

import (
	"os"
	"os/user"
	"path"
)

/**
 * @Note that it is advisable to not test if a path exists, as
 * that can cause race conditions, and can produce an invalid test
 * as the path could be created between the test, and the use of
 * the path.
 */

const (
	WUNDERTOOLS_PROJECT_CONF_FOLDER = ".wundertools"
	WUNDERTOOLS_USER_CONF_SUBPATH   = "wundertools"
)

// Struct used to keep path information
type Paths struct {
	allPaths     map[string]string // Paths is a string keyed map of important paths in the project
	confPathKeys []string          // ordered set of AllPaths keys that are possible roots for settings
}

// do some discovery of paths based on a pwd
func (paths *Paths) DefaultPaths(workingDir string) {
	paths.Init()
	paths.SetPath("working", workingDir, false)
	paths.DiscoverUserPaths()
	paths.DiscoverProjectPaths(workingDir)
}

// quick constructor/initializer
func (paths *Paths) Init() {
	paths.allPaths = make(map[string]string)
	paths.confPathKeys = make([]string, 0)
}

// a quick snippet to discover a user's home folder
func (paths *Paths) userHomePath() string {
	if existingPath, ok := paths.Path("user-home"); ok {
		return existingPath
	} else if currentUser, err := user.Current(); err == nil {
		return currentUser.HomeDir
	} else {
		return os.Getenv("HOME")
	}
}

// set a path using a key and path, optionally defined it as a configuration root
func (paths *Paths) SetPath(key string, path string, isConf bool) {
	paths.allPaths[key] = path
	if isConf {
		paths.confPathKeys = append(paths.confPathKeys, key)
	}
}

// get a path from a key
func (paths *Paths) Path(key string) (path string, found bool) {
	path, found = paths.allPaths[key]
	return
}

// get all of the conf path keys in order
func (paths *Paths) OrderedConfPathKeys() []string {
	return paths.confPathKeys
}

/**
 * Discover some user paths
 *
 * @NOTE we have to play some games for different OSes here
 *
 * dependening on OS, determine if the user has any settings
 * if so, add a conf path for them.
 */
func (paths *Paths) DiscoverUserPaths() {
	homeDir := paths.userHomePath()
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
	paths.SetPath("user-home", homeDir, false)
	paths.SetPath("user-wundertools", homeWTDir, true)
}

/**
 * Discover project paths
 *
 * Recursively navigate up the file path until we discover a folder that
 * has the key configuration subfolder in it.  That path is marked as the
 * application root, and the subfolder is marked as a conf path
 */
func (paths *Paths) DiscoverProjectPaths(workingDir string) {
	homeDir := paths.userHomePath()

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
	paths.SetPath("project-root", projectRootDirectory, false)
	paths.SetPath("project-wundertools", path.Join(projectRootDirectory, WUNDERTOOLS_PROJECT_CONF_FOLDER), true)
}
