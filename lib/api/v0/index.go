package api_v0

import (
	"os"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context, ewtHome string) IndexResponse {

	//get a listing of directories in the ewtHome path
	dirs, err := os.ReadDir(ewtHome)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	dirNames := []string{}
	for _, dir := range dirs {
		if dir.IsDir() {
			dirNames = append(dirNames, dir.Name())
		}
	}

	response := IndexResponse{
		Message:     "welcome to ewt-web API v0",
		EwtHome:     ewtHome,
		Directories: dirNames,
	}

	return response
}
