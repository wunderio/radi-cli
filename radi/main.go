package main

import (
	"context"
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"

	api_command "github.com/wunderkraut/radi-api/operation/command"
	cli_local "github.com/wunderkraut/radi-cli/local"
	"github.com/wunderkraut/radi-cli/version"
	handler_local "github.com/wunderkraut/radi-handlers/local"
)

var (
	debug       bool                           = false                            // default to disable debugging output
	internal    bool                           = false                            // use and display components that are considered internal
	workingDir  string                         = ""                               // can't use os.Cwd which returns multi-value
	environment string                         = "local"                          // default to using a local environment
	flags       []string                       = os.Args                          // store the cli args before they get used
	settings    handler_local.LocalAPISettings = handler_local.LocalAPISettings{} // just make sure that we have an object
	ctx         context.Context                = context.Background()             // this would allow us to terminate/timeout operations
)

func init() {

	/**
	 * early process of global flags using the core flag library
	 * instead of the cli flags functionality, as we need the
	 * global flags acted on before we go to add operations.
	 *
	 * @NOTE this may affect os.Args.
	 */

	/**
	 * PreProcess the "environment" flag
	 *
	 *   This needs to be preprocessed, because it's value
	 *   will affect what API elements are loaded, and
	 *   therefore what operations will be available.  This
	 *   means that we need it processed before the cli is
	 *   processed.
	 */
	flag.StringVar(&environment, "environment", environment, "Activate a specific environment")

	/**
	 * PreProcess the "debug" flag
	 *
	 * While not strictly necessary, we preprocess debug
	 * in order to get the logrus verbosity raised as early
	 * as possible.
	 */
	flag.BoolVar(&debug, "debug", debug, "Enable verbose debugging output")

	/**
	 * PreProcess the "internal" flag
	 *
	 * While not strictly necessary, we preprocess debug
	 * in order to get the logrus verbosity raised as early
	 * as possible.
	 */
	flag.BoolVar(&internal, "internal", internal, "Enable API components that are considered internal")

	/**
	 * The following flags are all flags that we use pass through
	 * to the cli, and ignore here, but we need them declared so
	 * that the flag library doesn't fial validation.
	 */
	flag.Bool("help", false, "show help")
	flag.Bool("h", false, "show help")
	flag.Bool("version", false, "show version")
	flag.Bool("v", false, "show version")

	// parse those preprocessed flags
	flag.Parse()

	// If the debug flags was set, then assign a log verbosity to the logrus lib
	if debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Enabling verbose debug output")
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if internal {
		log.Info("CLI: Showing internal components")
	}

}

func main() {

	/**
	 * Now we use the urfave/cli app to build out CLI application
	 * repeating the global flags, and setting up operations based
	 * on the current project configuration.
	 */
	app := &cli.App{}
	app.Name = "radi-cli"
	app.Usage = "Command line interface for Radi API."
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"
	app.Authors = []*cli.Author{&cli.Author{Name: "Wunder.IO", Email: "https://github.com/wunderkraut/radi-cli"}}

	/**
	 * We PrePprocess these global flags in init() in order
	 * to allow the debug and environment to be processed before
	 * we add operations, however we still add them as global
	 * flags to the cli app in order to get the UI out of it.
	 *
	 * This is a bit hackish, but not flawed in approach.
	 */
	app.Flags = []cli.Flag{
		cli.Flag(&cli.StringFlag{
			Name:        "environment",
			Usage:       "Activate a specific environment",
			Hidden:      false,
			Destination: &environment,
		}),
		cli.Flag(&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Enable verbose debugging output",
			EnvVars:     []string{"RADI_DEBUG"},
			Hidden:      false,
			Destination: &debug,
		}),
		cli.Flag(&cli.BoolFlag{
			Name:        "internal",
			Usage:       "Enable API components that are considered internal",
			EnvVars:     []string{"RADI_INTERNAL"},
			Hidden:      false,
			Destination: &internal,
		}),
	}

	/**
	 * Create a settings object, which will be used to create an
	 * API instance for run-time use.
	 */

	workingDir, _ = os.Getwd()
	settings = MakeLocalAPISettings(workingDir, ctx)

	// Discover the current User
	DiscoverCurrentUser(&settings)

	// Discover paths for the user like ~ and ~/.config/wundertools
	DiscoverUserPaths(&settings)
	DiscoverProjectPaths(&settings)

	// if we have an environment set, then discover it.
	if environment != "" {
		DiscoverEnvironmentPath(&settings, environment)
		log.WithFields(log.Fields{"environment": environment, "config-paths": settings.ConfigPaths}).Debug("Enabled specific environment")
	}

	// Quick settings debug
	log.WithFields(log.Fields{"settings": settings, "paths": settings.ConfigPaths}).Debug("Discovered Settings [final]")

	/**
	 * The settings object has been created, now create an API,
	 * and use it to get a list of operations, which we will
	 * convert to CLI commands.
	 */

	// Build a local API implementation from the settings
	local, _ := cli_local.MakeLocalAPI(settings)

	// Get a list of operations from the API
	localOps := local.Operations()

	// Add any "external" operations from the api to the app
	AppApiOperations(app, localOps, internal)

	// Add any commands from the api CommandWrapper to the app
	AppWrapperCommands(app, api_command.New_SimpleCommandWrapper(localOps), internal)

	// Run the App initializer again to process the added operations
	app.Setup()

	// Run the CLI command based on passed args
	app.Run(flags)
}
