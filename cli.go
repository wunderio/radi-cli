package main

import (
	"fmt"
	"os"

	"github.com/james-nesbitt/wundertools-go/config"
)

var (
	app config.Application
)

func init() {

	workingDir, _ := os.Getwd()
	app = *config.DefaultApplication(workingDir)

}

func main() {

	fmt.Println("--SETTINGS--")
	fmt.Println("Name:", app.Name)
	fmt.Println("Author:", app.Author)

	fmt.Println("--PATHS--")
	fmt.Println("Conf Path keys:", app.Paths.OrderedConfPathKeys())
	fmt.Println("All Paths:", app.Paths)

}
