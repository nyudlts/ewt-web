package router

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nyudlts/ewt-web/controllers"
	api_v0 "github.com/nyudlts/ewt-web/lib"
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
	addRoutes(engine, config)
}

func addRoutes(engine *gin.Engine, ewtConfig *EWTConfig) {
	//Index
	engine.GET("/", func(c *gin.Context) {
		indexResponse := api_v0.IndexHandler(c, ewtConfig.EWTHome)

		c.HTML(200, "index.html.tmpl", gin.H{
			"title": "Welcome to ewt-web",
			"data":  indexResponse,
		})
	})

	//archives
	archivesRoutes := engine.Group("archives")

	//show all archives
	archivesRoutes.GET("/", func(c *gin.Context) {
		archivesLocations := filepath.Join(ewtConfig.EWTHome, "completed")
		archives, err := api_v0.ListArchives(archivesLocations)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.HTML(200, "show_archives.html.tmpl", gin.H{
			"archives": archives,
		})
	})

	//download an archive tar
	archivesRoutes.GET("/:archive", func(c *gin.Context) {
		archiveName := c.Param("archive")
		api_v0.DownloadArchive(c, ewtConfig.EWTHome, archiveName)
	})

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

	projectRoutes.GET("new", func(c *gin.Context) { controllers.NewProject(c, ewtConfig.EWTHome) })

	projectRoutes.POST("", func(c *gin.Context) {
		controllers.CreateProject(c, ewtConfig.EWTHome)
	})

	projectRoutes.GET(":id/jobs/project", func(c *gin.Context) {
		projectName := c.Param("id")
		controllers.GetProjectJobs(c, ewtConfig.EWTHome, projectName)
	})

	projectRoutes.GET(":id/jobs/:job", func(c *gin.Context) {
		projectName := c.Param("id")
		job := c.Param("job")
		if result, err := api_v0.RouteProjectJob(ewtConfig.EWTHome, projectName, job); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			log.Println(result)
			c.Redirect(302, "/")
		}
	})

}
