package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lankaiyun/company/internal/controllers"
)

func RouterInit(r *gin.Engine) {
	mainRouter := r.Group("/")
	{
		mainRouter.GET("/", controllers.Controller{}.Index)
	}
}
