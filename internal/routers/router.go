package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lankaiyun/company/internal/controllers"
)

func RouterInit(r *gin.Engine) {
	showRouter := r.Group("/")
	{
		showRouter.GET("/", controllers.Controller{}.Home)
	}
}
