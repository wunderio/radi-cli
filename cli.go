package main

import (
	"fmt"
	"os"

	"github.com/james-nesbitt/wundertools-go/config"
)

var (
  paths *config.Paths
)

func init() {

	workingDir, _ := os.Getwd()
	paths = config.DefaultPaths(workingDir)

}

func main() {

	fmt.Println(paths)

}
