package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api_v0 "github.com/nyudlts/ewt-web/lib/api/v0"
)

func ShowProject(c *gin.Context, ewtRoot string, projectName string) {
	response, err := api_v0.ShowProjectHandler(c, ewtRoot, projectName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "project_show.html.tmpl", gin.H{
		"project": response,
	})
}

func GetProjectLogs(c *gin.Context, ewtRoot string, projectName string) {
	logFiles, err := api_v0.GetProjectLogsHandler(c, ewtRoot, projectName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "project_get_logs.html.tmpl", gin.H{
		"logs":        logFiles,
		"projectName": projectName,
	})
}

func ShowProjectLog(c *gin.Context, ewtRoot string, projectName string, logFileName string) {
	logContent, err := api_v0.ShowProjectLogHandler(c, ewtRoot, projectName, logFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "project_show_log.html.tmpl", gin.H{
		"projectName": projectName,
		"logFileName": logFileName,
		"logContent":  logContent,
	})
}
