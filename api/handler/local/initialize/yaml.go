package initialize

import (
	"io/ioutil"
	"net/http"
	"strings"

	"encoding/json"
	"gopkg.in/yaml.v2"

	log "github.com/Sirupsen/logrus"
)

// Get tasks from remote YAML corresponding to a remote yaml file
func (tasks *InitTasks) Init_Yaml_Run(path string) bool {

	var yamlSourceBytes []byte
	var err error

	if strings.Contains(path, "://") {

		resp, err := http.Get(path)
		if err != nil {
			log.WithFields(log.Fields{"path": path}).WithError(err).Error("Could not retrieve remote yaml init instructions.")
			return false
		}
		defer resp.Body.Close()
		yamlSourceBytes, err = ioutil.ReadAll(resp.Body)

	} else {

		// read the config file
		yamlSourceBytes, err = ioutil.ReadFile(path)
		if err != nil {
			log.WithFields(log.Fields{"path": path}).WithError(err).Error("Could not read the local YAML file.")
			return false
		}
		if len(yamlSourceBytes) == 0 {
			log.WithFields(log.Fields{"path": path}).Error("Yaml file was empty.")
			return false
		}

	}

	tasks.AddMessage("Initializing using YAML Source [" + path + "] to local project folder")

	// get tasks from yaml
	if err := tasks.AddTasksFromYaml(yamlSourceBytes); err != nil {
		log.WithError(err).Error("An error occured interpreting yml task list.")
	}

	// Add some message items
	tasks.AddFile("kraut/CREATEDFROM.md", "THIS PROJECT WAS CREATED A KRAUT YAML INSTALLER :"+path)

	return true
}

/**
 *Getting tasks from YAML
 */

func (tasks *InitTasks) AddTasksFromYaml(yamlSource []byte) error {

	var yaml_tasks []map[string]interface{}
	err := yaml.Unmarshal(yamlSource, &yaml_tasks)
	if err != nil {
		return err
	}

	var taskAdder TaskAdder
	for _, task_struct := range yaml_tasks {

		taskAdder = nil

		if _, ok := task_struct["Type"]; !ok {
			continue
		}

		switch task_struct["Type"] {
		case "File":
			json_task, _ := json.Marshal(task_struct)
			var task InitTaskYaml_FileMake
			if err := json.Unmarshal(json_task, &task); err == nil {
				taskAdder = TaskAdder(&task)
			}
		case "RemoteFile":
			json_task, _ := json.Marshal(task_struct)
			var task InitTaskYaml_RemoteFileCopy
			if err := json.Unmarshal(json_task, &task); err == nil {
				taskAdder = TaskAdder(&task)
			}
		case "FileCopy":
			json_task, _ := json.Marshal(task_struct)
			var task InitTaskYaml_FileCopy
			if err := json.Unmarshal(json_task, &task); err == nil {
				taskAdder = TaskAdder(&task)
			}
		case "FileStringReplace":
			json_task, _ := json.Marshal(task_struct)
			var task InitTaskYaml_FileStringReplace
			if err := json.Unmarshal(json_task, &task); err == nil {
				taskAdder = TaskAdder(&task)
			}
		case "GitClone":
			json_task, _ := json.Marshal(task_struct)
			var task InitTaskYaml_GitClone
			if err := json.Unmarshal(json_task, &task); err == nil {
				taskAdder = TaskAdder(&task)
			}
		case "Message":
			json_task, _ := json.Marshal(task_struct)
			var task InitTaskYaml_Message
			if err := json.Unmarshal(json_task, &task); err == nil {
				taskAdder = TaskAdder(&task)
			}
		case "Error":
			json_task, _ := json.Marshal(task_struct)
			var task InitTaskYaml_Error
			if err := json.Unmarshal(json_task, &task); err == nil {
				taskAdder = TaskAdder(&task)
			}

		default:
			taskType := task_struct["Type"].(string)
			log.WithFields(log.Fields{"type": taskType}).Warning("Unknown init task type.")
		}

		if taskAdder != nil {
			taskAdder.AddTask(tasks)
		}

	}

	return nil
}

type InitTaskYaml_Base struct {
	Type string `json:"Type" yaml:"Type"`
}

type TaskAdder interface {
	AddTask(tasks *InitTasks)
}

type InitTaskYaml_FileMake struct {
	Path     string `json:"Path" yaml:"Path"`
	Contents string `json:"Contents" yaml:"Contents"`
}

func (task *InitTaskYaml_FileMake) AddTask(tasks *InitTasks) {
	tasks.AddFile(task.Path, task.Contents)
}

type InitTaskYaml_RemoteFileCopy struct {
	Path string `json:"Path" yaml:"Path"`
	Url  string `json:"Url" yaml:"Url"`
}

func (task *InitTaskYaml_RemoteFileCopy) AddTask(tasks *InitTasks) {
	tasks.AddRemoteFile(task.Path, task.Url)
}

type InitTaskYaml_FileCopy struct {
	Path   string `json:"Path" yaml:"Path"`
	Source string `json:"Source" yaml:"Source"`
}

func (task *InitTaskYaml_FileCopy) AddTask(tasks *InitTasks) {
	tasks.AddFileCopy(task.Path, task.Source)
}

type InitTaskYaml_FileStringReplace struct {
	Path         string `json:"Path" yaml:"Path"`
	OldString    string `json:"Old" yaml:"Source"`
	NewString    string `json:"New" yaml:"Source"`
	ReplaceCount int    `json:"Limit" yaml:"Source"`
}

func (task *InitTaskYaml_FileStringReplace) AddTask(tasks *InitTasks) {
	tasks.AddFileStringReplace(task.Path, task.OldString, task.NewString, task.ReplaceCount)
}

type InitTaskYaml_GitClone struct {
	Path string `json:"Path" yaml:"Path"`
	Url  string `json:"Url" yaml:"Url"`
}

func (task *InitTaskYaml_GitClone) AddTask(tasks *InitTasks) {
	tasks.AddGitClone(task.Path, task.Url)
}

type InitTaskYaml_Message struct {
	Message string `json:"Message" yaml:"Message"`
}

func (task *InitTaskYaml_Message) AddTask(tasks *InitTasks) {
	tasks.AddMessage(task.Message)
}

type InitTaskYaml_Error struct {
	Error string `json:"Error" yaml:"Error"`
}

func (task *InitTaskYaml_Error) AddTask(tasks *InitTasks) {
	tasks.AddError(task.Error)
}
