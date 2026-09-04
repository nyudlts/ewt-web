package controllers

import "github.com/gin-gonic/gin"

func GetIndex(c *gin.Context, indexResponse interface{}) {
	c.HTML(200, "index.html.tmpl", gin.H{
		"title": "Welcome to ewt-web",
		"data":  indexResponse,
	})
}
