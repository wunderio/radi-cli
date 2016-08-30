package compose

/**
 * Log some information about the compose nodes
 */

import (
	"strings"
	// "strconv"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	// "github.com/james-nesbitt/wundertools-go/config"
	// "github.com/docker/libcompose/docker"
	libCompose_project "github.com/docker/libcompose/project"
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
		infoSet, _ := service.Info(context.Background())
		containers, _ := service.Containers(ctx)
		config := service.Config()

		var info libCompose_project.Info
		if len(infoSet) > 0 {
			info = infoSet[0]
		} else {
			info = libCompose_project.Info{}
		}

		if _, ok := info["State"]; !ok {
			info["State"] = ""
		}
		// infoString := ""
		// for index, info := range infoSet {
		// 	for key, value := range info {
		// 		infoString += strconv.FormatInt(int64(index), 10)+":"+key+"="+value+"|"
		// 	}
		// }

		var row []string

		row = []string{
			"|-",
			service.Name(),
			config.Image,
			"",
			info["State"],
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
