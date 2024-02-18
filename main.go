package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lankaiyun/company/global"
	"github.com/lankaiyun/company/internal/routers"
)

func main() {
	global.Logger = InitLog()
	gin.SetMode(gin.ReleaseMode)
	router := NewRouter()
	err := router.Run(":7777")
	if err != nil {
		global.Logger.SetPrefix("[Fatal] ")
		global.Logger.Fatalf("router.Run reported an error! err: %v\n", err)
		return
	}
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
	routers.RouterInit(r)
	return r
}

func InitLog() *log.Logger {
	f, err := os.OpenFile("company.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	return log.New(f, "", log.LstdFlags|log.Llongfile)
}
