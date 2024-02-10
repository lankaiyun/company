package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lankaiyun/company/internal/routers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := NewRouter()
	router.Run(":7777")
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// static resource
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
	routers.RouterInit(r)
	return r
}
