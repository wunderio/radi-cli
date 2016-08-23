package command

import (
	"os/exec"

	"github.com/james-nesbitt/wundertools-go/config"

	log "github.com/Sirupsen/logrus"
)

const COMMAND_TYPE_EXEC = "exec"

type ExecCommandSettings struct {
	EnvironmentVars map[string]string `yaml:"vars"`
	Exec            []string          `yaml:"exec"`
	RunDir          string            `yaml:"path"`
}

type ExecCommand struct {
	CommandBase
	settings   ExecCommandSettings
	persistant bool
}

func (command *ExecCommand) Init(application *config.Application) {
	command.CommandBase.Init(application)
}
func (command *ExecCommand) Settings(settings interface{}) {
	command.settings = settings.(ExecCommandSettings)
}
func (command *ExecCommand) Exec(flags ...string) {

	writer := log.StandardLogger().Writer()
	defer writer.Close()

	execCmd := ""
	execArgs := []string{}

	switch len(command.settings.Exec) {
	case 0:
		log.Error("Not enough arguments defined in the exec command")
		return
	case 1:
		execCmd = command.settings.Exec[0]
	default:
		execCmd = command.settings.Exec[0]
		execArgs = command.settings.Exec[1:]
	}

	cmd := exec.Command(execCmd, execArgs...)
	// cmd.Stdin = strings.NewReader("some input")
	cmd.Stdout = writer

	if len(command.settings.EnvironmentVars) > 0 {
		for key, value := range command.settings.EnvironmentVars {
			cmd.Env = append(cmd.Env, key+"="+value)
		}
	}
	if command.settings.RunDir != "" {
		if systemPath, exists := command.application.Paths.Path(command.settings.RunDir); exists {
			cmd.Dir = systemPath
		} else {
			cmd.Dir = command.settings.RunDir
		}
	} else {
		cmd.Dir, _ = command.application.Paths.Path("project-root")
	}

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
