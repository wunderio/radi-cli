package initialize

import (
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

func (tasks *InitTasks) Init_Git_Run(source string) bool {

	if source == "" {
		log.Error("You have not provided a git target $/> wundertools init git https://github.com/aleksijohansson/docker-drupal-coach")
		return false
	}

	url := source
	path := tasks.root

	logWriter := log.StandardLogger().Writer()
	defer logWriter.Close()

	cmd := exec.Command("git", "clone", "--progress", url, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = logWriter
	cmd.Stderr = logWriter

	err := cmd.Start()

	if err != nil {
		log.WithFields(log.Fields{"url": url}).WithError(err).Error("Failed to clone the remote repository.")
		return false
	}

	log.WithFields(log.Fields{"url": url}).Info("Clone remote repository to local project folder.")
	err = cmd.Wait()

	if err != nil {
		log.WithFields(log.Fields{"url": url}).WithError(err).Error("Failed to clone the remote repository.")
		return false
	}

	tasks.AddMessage("Cloned remote repository [" + url + "] to local project folder")
	tasks.AddFile(".wundertools/CREATEDFROM.md", `THIS PROJECT WAS CREATED FROM GIT`)

	return true
}

type InitTaskGitClone struct {
	InitTaskFileBase
	root string

	path string
	url  string
}

func (task *InitTaskGitClone) RunTask() bool {
	if task.root == "" || task.url == "" {
		log.WithFields(log.Fields{"root": task.root}).Error("EMPTY ROOT PASSED TO GIT")
		return false
	}

	destinationPath := task.path
	url := task.url

	logWriter := log.StandardLogger().Writer()
	defer logWriter.Close()

	if !task.MakeDir(destinationPath, false) {
		return false
	}

	destinationAbsPath, ok := task.absolutePath(destinationPath, true)
	if !ok {
		log.WithFields(log.Fields{"path": destinationPath}).Warning("Invalid copy destination path.")
		return false
	}

	cmd := exec.Command("git", "clone", "--progress", url, destinationAbsPath)
	cmd.Stderr = logWriter
	err := cmd.Start()

	if err != nil {
		log.WithFields(log.Fields{"url": url}).WithError(err).Error("Failed to clone the remote repository.")
		return false
	}

	err = cmd.Wait()

	if err != nil {
		log.WithFields(log.Fields{"url": url}).WithError(err).Error("Failed to clone the remote repository.")
		return false
	}

	log.WithFields(log.Fields{"url": url, "path": destinationPath}).Info("Cloned remote repository to local path.")
	return true
}
