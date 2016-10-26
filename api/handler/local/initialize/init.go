package initialize

import (
	log "github.com/Sirupsen/logrus"
)

type InitTasks struct {
	root string

	tasks []InitTask
}

func (tasks *InitTasks) Init(root string) bool {
	tasks.root = root
	tasks.tasks = []InitTask{}
	return true
}
func (tasks *InitTasks) RunTasks() {
	for _, task := range tasks.tasks {
		log.WithFields(log.Fields{"task": task}).Debug("INIT TASK")
		task.RunTask()
	}
}

func (tasks *InitTasks) AddTask(task InitTask) {
	tasks.tasks = append(tasks.tasks, task)
}

func (tasks *InitTasks) AddFile(path string, contents string) {
	tasks.AddTask(InitTask(&InitTaskFile{
		root:     tasks.root,
		path:     path,
		contents: contents,
	}))
}
func (tasks *InitTasks) AddRemoteFile(path string, url string) {
	tasks.AddTask(InitTask(&InitTaskRemoteFile{
		root: tasks.root,
		path: path,
		url:  url,
	}))
}
func (tasks *InitTasks) AddFileCopy(path string, source string) {
	tasks.AddTask(InitTask(&InitTaskFileCopy{
		root:   tasks.root,
		path:   path,
		source: source,
	}))
}
func (tasks *InitTasks) AddFileStringReplace(path string, oldString string, newString string, replaceCount int) {
	tasks.AddTask(InitTask(&InitTaskFileStringReplace{
		root:         tasks.root,
		path:         path,
		oldString:    oldString,
		newString:    newString,
		replaceCount: replaceCount,
	}))
}
func (tasks *InitTasks) AddGitClone(path string, url string) {
	tasks.AddTask(InitTask(&InitTaskGitClone{
		root: tasks.root,
		path: path,
		url:  url,
	}))
}
func (tasks *InitTasks) AddMessage(message string) {
	tasks.AddTask(InitTask(&InitTaskMessage{
		message: message,
	}))
}
func (tasks *InitTasks) AddError(error string) {
	tasks.AddTask(InitTask(&InitTaskError{
		error: error,
	}))
}

type InitTask interface {
	RunTask() bool
}

type InitTaskError struct {
	error string
}

func (task *InitTaskError) RunTask() bool {
	log.WithFields(log.Fields{"task": task}).Error(task.error)
	return true
}

type InitTaskMessage struct {
	message string
}

func (task *InitTaskMessage) RunTask() bool {
	log.WithFields(log.Fields{"task": task}).Info(task.message)
	return true
}
