package router

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nyudlts/ewt-web/controllers"
)

type EWTConfig struct {
	EWTHome string `json:"ewt-home"`
}

func InitRouter(engine *gin.Engine, config *EWTConfig) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	templatesPath := filepath.Join(wd, "templates/**/*.tmpl")
	engine.LoadHTMLGlob(templatesPath)
	addAPIRoutes(engine, config)
}

func addAPIRoutes(engine *gin.Engine, ewtConfig *EWTConfig) {
	//Index
	engine.GET("/", func(c *gin.Context) { controllers.GetIndex(c, ewtConfig.EWTHome) })

	//Projects
	projectRoutes := engine.Group("projects")
	projectRoutes.GET(":id/show", func(c *gin.Context) {
		projectName := c.Param("id")
		controllers.ShowProject(c, ewtConfig.EWTHome, projectName)
	})

	projectRoutes.GET(":id/logs", func(c *gin.Context) {
		projectName := c.Param("id")
		controllers.GetProjectLogs(c, ewtConfig.EWTHome, projectName)
	})

	projectRoutes.GET(":id/logs/:log", func(c *gin.Context) {
		projectName := c.Param("id")
		logFileName := c.Param("log")
		controllers.ShowProjectLog(c, ewtConfig.EWTHome, projectName, logFileName)
	})

}
