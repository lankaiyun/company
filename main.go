package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lankaiyun/company/global"
	"github.com/lankaiyun/company/internal/database"
	"github.com/lankaiyun/company/internal/routers"
	"github.com/lankaiyun/company/pkg/log"
	"github.com/lankaiyun/company/pkg/setting"
)

func init() {
	global.SugarLogger = log.GetLogger()
	err := SetupSetting()
	if err != nil {
		global.SugarLogger.Error("SetupSetting err:", err)
	}
	global.MySqlConn = database.GetDbObj(global.DatabaseSetting)
}

func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.SugarLogger.Error("ListenAndServer err:", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.SugarLogger.Fatal("Shut down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		global.SugarLogger.Fatal("Server forced to shutdown:", err)
	}
	global.SugarLogger.Info("Server Exiting")
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "main/404.html", nil)
	})
	routers.RouterInit(r)
	return r
}
