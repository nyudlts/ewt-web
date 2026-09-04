package router

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nyudlts/ewt-web/controllers"
	api_v0 "github.com/nyudlts/ewt-web/lib/api/v0"
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

func addAPIRoutes(engine *gin.Engine, ewtRoot *EWTConfig) {
	engine.GET("/", func(c *gin.Context) {
		indexResponse := api_v0.IndexHandler(c, ewtRoot.EWTHome)
		controllers.GetIndex(c, indexResponse)
	})
}
