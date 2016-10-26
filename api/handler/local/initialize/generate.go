package initialize

import (
	"bufio"
	"io"
	"os"
	"path"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func Init_Generate(handler string, path string, skip []string, sizeLimit int64, output io.Writer) bool {
	log.Info("GENERATING INIT")

	var generator Generator
	switch handler {
	case "test":
		generator = Generator(&TestInitGenerator{output: output})
	case "yaml":
		generator = Generator(&YMLInitGenerator{output: output})
	default:
		log.WithFields(log.Fields{"handler": handler}).Error("Unknown init generator (handler).")
		return false
	}

	iterator := GenerateIterator{
		output:    output,
		skipFiles: []string{".gitignore", ".ignore"}, // @todo get this from function parameter?
		skip:      skip,
		sizeLimit: sizeLimit,
		generator: generator,
	}

	if iterator.Generate(path) {
		log.Info("FINISHED GENERATING YML INIT")
		return true
	} else {
		log.Error("ERROR OCCURRED GENERATING YML INIT")
		return false
	}
}

type GenerateIterator struct {
	output io.Writer

	skipFiles []string
	skip      []string
	sizeLimit int64

	generator Generator
}

func (iterator *GenerateIterator) Generate(path string) bool {
	return iterator.generate_Recursive(path, "", iterator.skip)
}
func (iterator *GenerateIterator) generate_Recursive(sourceRootPath string, sourcePath string, skip []string) bool {
	fullPath := sourceRootPath

	if sourcePath != "" {
		fullPath = path.Join(fullPath, sourcePath)
	}

	for _, skipEach := range skip {
		if match, _ := regexp.MatchString(skipEach, sourcePath); match {
			log.WithFields(log.Fields{"path": sourcePath}).Info("Skipping marked skip item.")
			return true
		}
	}

	// get properties of source dir
	info, err := os.Stat(fullPath)
	if err != nil {
		// @TODO do something log : source doesn't exist
		log.WithFields(log.Fields{"path": fullPath}).Warning("File does not exist.")
		return false
	}

	mode := info.Mode()
	if mode.IsDir() {

		// check for GIT folder
		if _, err := os.Open(path.Join(fullPath, ".git")); err == nil {
			if iterator.generator.generateGit(fullPath, sourcePath) {
				log.WithFields(log.Fields{"path": sourcePath}).Info("Generated git file.")
				return true
			} else {
				log.WithFields(log.Fields{"path": sourcePath}).Warning("Failed to generate git file.")
			}
		}

		// add any .gitignore/.dockerignore entries to the skip list
		for _, ignoreFileName := range iterator.skipFiles {
			if ignoreFile, err := os.Open(path.Join(fullPath, ignoreFileName)); err == nil {

				scanner := bufio.NewScanner(ignoreFile)
				for scanner.Scan() {
					text := scanner.Text()

					if text == "" {
						continue
					}
					if strings.HasPrefix(text, "#") {
						continue
					}

					log.WithFields(log.Fields{"skip": text, "file": ignoreFileName, "path": fullPath}).Info("add ignore item to skip list")
					skip = append(skip, strings.TrimSpace(text))
				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}

			}
		}

		directory, _ := os.Open(fullPath)
		defer directory.Close()
		objects, err := directory.Readdir(-1)

		if err != nil {
			// @TODO do something log : source doesn't exist
			log.WithFields(log.Fields{"path": fullPath}).Warning("Could not open directory.")
			return false
		}

		for _, obj := range objects {

			//childSourcePath := source + "/" + obj.Name()
			childSourcePath := path.Join(sourcePath, obj.Name())
			if !iterator.generate_Recursive(sourceRootPath, childSourcePath, skip) {
				log.WithFields(log.Fields{"path": childSourcePath, "root": sourceRootPath}).Warning("Resursive generate failed")
			}

		}

	} else if mode.IsRegular() {

		if info.Size() > iterator.sizeLimit {
			log.WithFields(log.Fields{"path": sourcePath, "limit": iterator.sizeLimit}).Info("Skipped file that is larger than our limit.")
			return true
		}

		// generate single file from contents
		if iterator.generator.generateSingleFile(fullPath, sourcePath) {
			log.WithFields(log.Fields{"path": sourcePath}).Info("Generated file (recursively).")
			return true
		} else {
			log.WithFields(log.Fields{"path": sourcePath}).Warning("Failed to generate file.")
			return false
		}
		return true
	} else {
		log.WithFields(log.Fields{"path": sourcePath}).Warning("Skipped generation non-regular file.")
	}

	return true
}

type Generator interface {
	generateSingleFile(fullPath string, sourcePath string) bool
	generateGit(fullPath string, sourcePath string) bool
}

type TestInitGenerator struct {
	output io.Writer
}

func (generator *TestInitGenerator) generateSingleFile(fullPath string, sourcePath string) bool {
	singleFile, _ := os.Open(fullPath)
	defer singleFile.Close()

	log.WithFields(log.Fields{"name": singleFile.Name()}).Debug("GENERATE SINGLE FILE")
	generator.output.Write([]byte("GENERATE SINGLE FILE: " + sourcePath + "\n"))
	return true
}
func (generator *TestInitGenerator) generateGit(fullPath string, sourcePath string) bool {
	log.WithFields(log.Fields{"path": sourcePath}).Debug("GENERATE GIT FILE")
	generator.output.Write([]byte("GENERATE GIT FILE: " + sourcePath + "\n"))
	return true
}
