package controllers

import (
	"github.com/gin-gonic/gin"
	api_v0 "github.com/nyudlts/ewt-web/lib/api/v0"
)

func GetIndex(c *gin.Context, ewtRoot string) {
	indexResponse := api_v0.IndexHandler(c, ewtRoot)

	c.HTML(200, "index.html.tmpl", gin.H{
		"title": "Welcome to ewt-web",
		"data":  indexResponse,
	})
}
