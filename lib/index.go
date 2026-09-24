package lib

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context, ewtHome string) IndexResponse {

	//get a listing of directories in the ewtHome path
	dirNames := GetProjects(ewtHome)

	response := IndexResponse{
		Message:     "welcome to ewt-web API v0",
		EwtHome:     ewtHome,
		Directories: dirNames,
	}

	return response
}
