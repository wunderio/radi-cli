package api

/**
 * This file provides a more advanced API handler for a project
 * which allows a stronger definition of which handlers are used
 * in a project, and which operations should come from which
 * handlers.
 */

// ProjectAPI provides an API that is picks handlers and operations based on configuration
type ProjectAPI struct {
}

// Define the settings needed to configure a project
type ProjectAPISetup struct {
	Handlers []ProjectAPISetupHandler `json:"handlers" yaml:"handlers"` // Handler configuration
}

// Define Handlers for a project
type ProjectAPISetupHandler struct {
	Settings   interface{} `json:"settings" yaml:"setting"`      // Settings, which are an unknown structure and passed to the handler
	Operations []string    `json:"operations" yaml:"operations"` // Which operations to use from this Handler
}
