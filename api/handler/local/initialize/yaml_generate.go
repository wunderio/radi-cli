package initialize

import (
	"bufio"
	"io"
	"os"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type YMLInitGenerator struct {
	output io.Writer
}

func (generator *YMLInitGenerator) generateSingleFile(fullPath string, sourcePath string) bool {
	singleFile, _ := os.Open(fullPath)
	defer singleFile.Close()

	log.WithFields(log.Fields{"name": singleFile.Name()}).Debug("GENERATE SINGLE FILE")
	generator.output.Write([]byte("- Type: File\n"))
	generator.output.Write([]byte("  path: " + sourcePath + "\n"))

	r := bufio.NewReader(singleFile)
	indent := "    "

	line, err := r.ReadString(10) // 0x0A separator = newline
	if err == io.EOF {
		generator.output.Write([]byte("  Contents: \n"))
	} else if err != nil {
		log.Error(err.Error())
		return false
	} else {

		generator.output.Write([]byte("  Contents: |2\n"))
		line = "" + line

		for {
			generator.output.Write([]byte(indent + line))

			line, err = r.ReadString(10) // 0x0A separator = newline
			if err == io.EOF {
				generator.output.Write([]byte(indent + "\n"))
				break
			} else if err != nil {
				log.Error(err.Error())
				return false
			}
		}

	}
	return true
}
func (generator *YMLInitGenerator) generateGit(fullPath string, sourcePath string) bool {

	gitUrl := ""

	if configFile, err := os.Open(path.Join(fullPath, ".git", "config")); err == nil {

		r := bufio.NewReader(configFile)
		for {
			line, err := r.ReadString(10) // 0x0A separator = newline
			if err == io.EOF {
				break
			} else if err != nil {
				log.WithError(err).Error("Could not read from git configuration file.")
				return false // if you return error
			}
			if strings.Contains(line, "url =") {
				lineSplit := strings.Split(line, "url =")
				gitUrl = strings.Trim(lineSplit[len(lineSplit)-1], " ")
				break
			}
		}

	} else {
		log.WithFields(log.Fields{"path": sourcePath}).Error("Could not open .git/config in path.")
		return false
	}

	if gitUrl == "" {
		log.Error("Could not determine GIT Url from .git/config")
		return false
	}

	log.WithFields(log.Fields{"path": sourcePath}).Debug("GENERATE GIT FILE.")
	generator.output.Write([]byte("- Type: GitClone\n"))
	generator.output.Write([]byte("  path: " + sourcePath + "\n"))
	generator.output.Write([]byte("  Url: " + gitUrl + "\n"))
	return true
}
