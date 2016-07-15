package compose

/**
 * Log some information about the compose nodes
 */

import (
	"strings"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	// "github.com/james-nesbitt/wundertools-go/config"
	// "github.com/docker/libcompose/docker"
	// libCompose_project "github.com/docker/libcompose/project"
)

// main info method
func (project *ComposeProject) Info() {

	log.Info("PROJECT COMPOSE INFORMATION")

	project.info_nodes()

}

// write information about nodes
func (project *ComposeProject) info_nodes() {

	// use a single run context for all of the operations
	ctx := context.Background()

	writer := log.StandardLogger().Writer()
	defer writer.Close()

	w := new(tabwriter.Writer)
	w.Init(writer, 15, 20, 2, ' ', 0)

	row := []string{
		"|=",
		"Name",
		"Image",
		"Container",
		"Status",
	}
	w.Write([]byte(strings.Join(row, "\t") + "\n"))

	for _, serviceName := range project.serviceNames() {

		service, _ := project.CreateService(serviceName)
		containers, _ := service.Containers(ctx)
		config := service.Config()

		var row []string

		row = []string{
			"|-",
			service.Name(),
			config.Image,
			// 	image.ID[:11],
			// 	strings.Join(image.RepoTags, ","),
			// 	strconv.FormatInt(image.Created, 10),
			// 	//			strconv.FormatInt(image.Size, 10),
			// 	//			strconv.FormatInt(image.VirtualSize, 10),
			// 	//			image.ParentID,
			// 	// 			strings.Join(image.RepoDigests, "\n"),
			// 	// 			strings.Join(image.Labels, "\n"),
		}
		w.Write([]byte(strings.Join(row, "\t") + "\n"))

		if len(containers) > 0 {
			for _, container := range containers {

				var status string = "not-running"
				if running, _ := container.IsRunning(ctx); running {
					status = "running"
				}

				row = []string{
					"|-",
					" ",
					" ",
					container.Name(),
					status,
				}
				w.Write([]byte(strings.Join(row, "\t") + "\n"))

			}
		}
	}

	w.Flush()
}
