package initialize

func (tasks *InitTasks) Init_Default_Run(source string) bool {

	switch source {
	case "bare":
		fallthrough
	default:
		return tasks.Init_Default_Bare()
	}

	return false
}

func (tasks *InitTasks) Init_Default_Bare() bool {

	tasks.AddFile(".wundertools/settings.yml", `# Wundertools project conf
Project: bare`)
	tasks.AddFile("docker-compose.yml", `# Project services
`)
	tasks.AddFile("app/README.md", `# Bare Project
## /.wundertools/

  Project wundertools configuration

## /app

  Project source-code and assets path

`)

	tasks.AddMessage("Created local project as a `bare` project")

	return true
}
