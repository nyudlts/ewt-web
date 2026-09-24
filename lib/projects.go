package lib

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	restricedDirectories = []string{"completed", "home", "lost", "quarantine", "templates", "test", "jobs"}
)

func GetProjects(ewtRoot string) []string {
	dirs, err := os.ReadDir(ewtRoot)
	if err != nil {
		return nil
	}
	var projectDirs []string
	for _, dir := range dirs {
		if dir.IsDir() {
			skip := false
			for _, restricted := range restricedDirectories {
				if dir.Name() == restricted {
					skip = true
					break
				}
			}
			if !skip {
				projectDirs = append(projectDirs, dir.Name())
			}
		}
	}
	return projectDirs
}

func ShowProjectHandler(c *gin.Context, ewtRoot string, projectName string) (*ProjectConfig, error) {
	configPath := filepath.Join(ewtRoot, projectName, "config.json")
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var projectConfig ProjectConfig
	err = json.Unmarshal(configBytes, &projectConfig)
	if err != nil {
		return nil, err
	}
	return &projectConfig, nil
}

func GetProjectLogsHandler(c *gin.Context, ewtRoot string, projectName string) ([]string, error) {
	logsPath := filepath.Join(ewtRoot, projectName, "logs")
	files, err := os.ReadDir(logsPath)
	if err != nil {
		return nil, err
	}
	var logFiles []string
	for _, file := range files {
		if !file.IsDir() {
			logFiles = append(logFiles, file.Name())
		}
	}
	return logFiles, nil
}

func ShowProjectLogHandler(c *gin.Context, ewtRoot string, projectName string, logFileName string) (string, error) {
	logFilePath := filepath.Join(ewtRoot, projectName, "logs", logFileName)
	logBytes, err := os.ReadFile(logFilePath)
	if err != nil {
		return "", err
	}
	return string(logBytes), nil
}
