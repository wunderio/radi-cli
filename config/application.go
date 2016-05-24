package config

func DefaultApplication(workingDir string) *Application {
	app := new(Application)
	app.Init()

	app.Paths.DefaultPaths(workingDir)
	app.from_ConfYaml()

	return app
}

// Configuration for an application
type Application struct {
	Name   string
	Author string

	Environment string

	*Paths
}

// quick initializer/constructor
func (app *Application) Init() {
	app.Name = ""
	app.Author = ""

	app.Environment = ""

	app.Paths = new(Paths)
}
