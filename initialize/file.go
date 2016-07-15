package initialize

import (
	"io"
	"os"
	"os/user"
	"path"
	"strings"

	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type InitTaskFileBase struct {
	root string
}

func (task *InitTaskFileBase) absolutePath(targetPath string, addRoot bool) (string, bool) {
	if strings.HasPrefix(targetPath, "~") {
		return path.Join(task.userHomePath(), targetPath[1:]), !addRoot

		// I am not sure how reliable this function is
		// you passed an absolute path, so I can't add the root
		// } else if path.isAbs(targetPath) {
		// 	return targetPath, !addRoot

		// you passed a relative path, and want me to add the root
	} else if addRoot {
		return path.Join(task.root, targetPath), true

		// you passed path and don't want the root added (but is it already abs?)
	} else if targetPath != "" {
		return targetPath, true

		// you passed an empty string, and don't want the root added?
	} else {
		return targetPath, false
	}
}
func (task *InitTaskFileBase) userHomePath() string {
	if currentUser, err := user.Current(); err == nil {
		return currentUser.HomeDir
	} else {
		return os.Getenv("HOME")
	}
}

func (task *InitTaskFileBase) MakeDir(makePath string, pathIsFile bool) bool {
	if makePath == "" {
		return true // it's already made
	}

	if pathDirectory, ok := task.absolutePath(makePath, true); !ok {
		log.WithFields(log.Fields{"path": pathDirectory}).Warning("Invalid directory path")
		return false
	}
	pathDirectory := path.Join(task.root, makePath)
	if pathIsFile {
		pathDirectory = path.Dir(pathDirectory)
	}

	if err := os.MkdirAll(pathDirectory, 0777); err != nil {
		// @todo something log
		return false
	}
	return true
}
func (task *InitTaskFileBase) MakeFile(destinationPath string, contents string) bool {
	if !task.MakeDir(destinationPath, true) {
		// @todo something log
		return false
	}

	if destinationPath, ok := task.absolutePath(destinationPath, true); !ok {
		log.WithFields(log.Fields{"path": destinationPath}).Warning("Invalid file destination path")
		return false
	}

	fileObject, err := os.Create(destinationPath)
	defer fileObject.Close()
	if err != nil {
		// @todo something log
		return false
	}
	if _, err := fileObject.WriteString(contents); err != nil {
		// @todo something log
		return false
	}

	return true
}

func (task *InitTaskFileBase) CopyFile(destinationPath string, sourcePath string) bool {
	if destinationPath == "" || sourcePath == "" {
		log.Warning("empty source or destination passed for copy")
		return false
	}

	sourcePath, ok := task.absolutePath(sourcePath, false)
	if !ok {
		log.WithFields(log.Fields{"path": sourcePath}).Warning("Invalid copy source path")
		return false
	}
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		log.WithFields(log.Fields{"path": sourcePath}).WithError(err).Warning("could not copy file as it does not exist")
		return false
	}
	defer sourceFile.Close()

	if !task.MakeDir(destinationPath, true) {
		// @todo something log
		log.WithFields(log.Fields{"path": destinationPath}).Warning("could not copy file as the path to the destination file could not be created")
		return false
	}
	destinationAbsPath, ok := task.absolutePath(destinationPath, true)
	if !ok {
		log.WithFields(log.Fields{"path": destinationPath}).Warning("Invalid copy destination path")
		return false
	}

	destinationFile, err := os.Open(destinationAbsPath)
	if err == nil {
		log.WithFields(log.Fields{"path": destinationPath}).Warning("could not copy file as it already exists.")
		destinationFile.Close()
		return false
	}

	destinationFile, err = os.Create(destinationAbsPath)
	if err != nil {
		log.WithFields(log.Fields{"path": destinationPath}).WithError(err).Warning("could not copy file as destination file could not be created.")
		return false
	}

	defer destinationFile.Close()
	_, err = io.Copy(destinationFile, sourceFile)

	if err == nil {
		sourceInfo, err := os.Stat(sourcePath)
		if err == nil {
			err = os.Chmod(destinationPath, sourceInfo.Mode())
			return true
		} else {
			log.WithFields(log.Fields{"path": destinationPath}).WithError(err).Warning("could not copy file as destination file could not be created.")
			return false
		}
	} else {
		log.WithFields(log.Fields{"path": destinationPath}).WithError(err).Warning("could not copy file as copy failed.")
	}

	return true
}

func (task *InitTaskFileBase) CopyRemoteFile(destinationPath string, sourcePath string) bool {
	if destinationPath == "" || sourcePath == "" {
		return false
	}

	response, err := http.Get(sourcePath)
	if err != nil {
		log.WithFields(log.Fields{"path": sourcePath}).Warning("Could not open remote URL")
		return false
	}
	defer response.Body.Close()

	sourceContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.WithFields(log.Fields{"path": sourcePath}).Warning("Could not read remote file")
		return false
	}

	return task.MakeFile(destinationPath, string(sourceContent))
}

func (task *InitTaskFileBase) CopyFileRecursive(path string, source string) bool {
	sourceAbsPath, ok := task.absolutePath(source, false)
	if !ok {
		log.WithFields(log.Fields{"path": source}).Warning("Couldn't find copy source.")
		return false
	}
	return task.copyFileRecursive(path, sourceAbsPath, "")
}
func (task *InitTaskFileBase) copyFileRecursive(destinationRootPath string, sourceRootPath string, sourcePath string) bool {
	fullPath := sourceRootPath

	if sourcePath != "" {
		fullPath = path.Join(fullPath, sourcePath)
	}

	// get properties of source dir
	info,
		err := os.Stat(fullPath)
	if err != nil {
		// @TODO do something log : source doesn't exist
		log.WithFields(log.Fields{"path": fullPath}).Warning("File does not exist.")
		return false
	}

	mode := info.Mode()
	if mode.IsDir() {

		directory, _ := os.Open(fullPath)
		objects, err := directory.Readdir(-1)

		if err != nil {
			// @TODO do something log : source doesn't exist
			log.WithFields(log.Fields{"path": fullPath}).Warning("Could not open directory")
			return false
		}

		for _, obj := range objects {

			//childSourcePath := source + "/" + obj.Name()
			childSourcePath := path.Join(sourcePath, obj.Name())
			if !task.copyFileRecursive(destinationRootPath, sourceRootPath, childSourcePath) {
				log.Warning("Resursive copy failed")
			}

		}

	} else {
		// add file copy
		destinationPath := path.Join(destinationRootPath, sourcePath)
		if task.CopyFile(destinationPath, sourceRootPath) {
			log.WithFields(log.Fields{"path": sourcePath, "root": sourceRootPath}).Info("--> Copied file (recursively).")
			return true
		} else {
			log.WithFields(log.Fields{"path": sourcePath, "root": sourceRootPath}).Warning("--> Failed to copy file.")
			return false
		}
		return true
	}
	return true
}

// perform a string replace on file contents
func (task *InitTaskFileBase) FileStringReplace(targetPath string, oldString string, newString string, replaceCount int) bool {

	targetPath, ok := task.absolutePath(targetPath, false)
	if !ok {
		log.WithFields(log.Fields{"path": targetPath}).Warning("Invalid string replace path")
		return false
	}

	contents, err := ioutil.ReadFile(targetPath)
	if err != nil {
		log.WithError(err).Error("Could not read file for string replacement.")
	}

	contents = []byte(strings.Replace(string(contents), oldString, newString, replaceCount))

	err = ioutil.WriteFile(targetPath, contents, 0644)
	if err != nil {
		log.WithError(err).Error("Could not write to file for string replacement.")
	}
	return true
}

type InitTaskFile struct {
	InitTaskFileBase
	root string

	path     string
	contents string
}

func (task *InitTaskFile) RunTask() bool {
	if task.path == "" {
		return false
	}

	if task.MakeFile(task.path, task.contents) {
		log.WithFields(log.Fields{"path": task.path}).Info("--> Created file.")
		return true
	} else {
		log.WithFields(log.Fields{"path": task.path}).Warning("--> Failed to create file.")
		return false
	}
}

type InitTaskRemoteFile struct {
	InitTaskFileBase
	root string

	path string
	url  string
}

func (task *InitTaskRemoteFile) RunTask() bool {
	if task.path == "" || task.root == "" || task.url == "" {
		return false
	}

	if task.CopyRemoteFile(task.path, task.url) {
		log.WithFields(log.Fields{"root": task.root, "url": task.url, "path": task.path}).Info("--> Copied remote file.")
		return true
	} else {
		log.WithFields(log.Fields{"root": task.root, "url": task.url, "path": task.path}).Warning("--> Failed to copy remote file.")
		return false
	}
}

type InitTaskFileCopy struct {
	InitTaskFileBase
	root string

	path   string
	source string
}

func (task *InitTaskFileCopy) RunTask() bool {
	if task.path == "" || task.root == "" || task.source == "" {
		return false
	}

	if task.CopyFileRecursive(task.path, task.source) {
		log.WithFields(log.Fields{"root": task.root, "path": task.path, "source": task.source}).Info("--> Copied file.")
		return true
	} else {
		log.WithFields(log.Fields{"root": task.root, "path": task.path, "source": task.source}).Warning("--> Failed to copy file.")
		return false
	}
}

type InitTaskFileStringReplace struct {
	InitTaskFileBase
	root string

	path         string
	oldString    string
	newString    string
	replaceCount int
}

func (task *InitTaskFileStringReplace) RunTask() bool {
	if task.path == "" || task.root == "" || task.oldString == "" || task.newString == "" {
		return false
	}
	if task.replaceCount == 0 {
		task.replaceCount = -1
	}

	if task.FileStringReplace(task.path, task.oldString, task.newString, task.replaceCount) {
		log.WithFields(log.Fields{"root": task.root, "path": task.path}).Info("--> performed string replace on file.")
		return true
	} else {
		log.WithFields(log.Fields{"root": task.root, "path": task.path}).Warning("--> Failed to perform string replace on file.")
		return false
	}
}
