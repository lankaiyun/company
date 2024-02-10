package controllers

import "github.com/gin-gonic/gin"

type Controller struct{}

func (ctl Controller) Home(c *gin.Context) {
	c.HTML(200, "main/index.html", nil)
}
