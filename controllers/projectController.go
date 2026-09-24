package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	ewt "github.com/nyudlts/electronic-records-workflow-tool/lib"
	api_v0 "github.com/nyudlts/ewt-web/lib"
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

func NewProject(c *gin.Context, ewtRoot string) {
	c.HTML(http.StatusOK, "project_new.html.tmpl", gin.H{
		"ewtRoot": ewtRoot,
	})
}

func CreateProject(c *gin.Context, ewtRoot string) {
	projectName := c.PostForm("project_name")
	sourcePath := c.PostForm("source_path")

	/* this needs to moved to jobs */
	configLocation := filepath.Join(ewtRoot, "ewt.json")
	if err := ewt.InitProject(projectName, sourcePath, configLocation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusSeeOther, "/projects/"+projectName+"/show")
}

func GetProjectJobs(c *gin.Context, ewtRoot string, projectName string) {
	projectJobs := api_v0.GetProjectJobs(ewtRoot, projectName)
	c.HTML(http.StatusOK, "get_jobs_project.html.tmpl", gin.H{
		"projectName": projectName,
		"jobs":        projectJobs,
	})
}
