package lib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ListArchives(ewtHome string) (map[string]int, error) {
	fmt.Println("Listing archives in:", ewtHome)
	files, err := os.ReadDir(ewtHome)
	if err != nil {
		return nil, err
	}
	archiveNames := map[string]int{}
	for _, file := range files {
		fmt.Println(file.Name())
		if !file.IsDir() {
			info, err := file.Info()
			if err != nil {
				continue
			}
			archiveNames[file.Name()] = int(info.Size())
		}
	}
	return archiveNames, nil
}

func DownloadArchive(c *gin.Context, ewtHome string, archiveName string) {
	archivesLocations := filepath.Join(ewtHome, "completed")
	archivePath := filepath.Join(archivesLocations, archiveName)
	c.FileAttachment(archivePath, archiveName)
}
