package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lankaiyun/company/internal/controllers"
)

func RouterInit(r *gin.Engine) {
	router := r.Group("/")
	{
		router.GET("/", controllers.Controller{}.Index)
		router.GET("/contact", controllers.Controller{}.Contact)

		router.POST("/addSubscription", controllers.Controller{}.AddSubscription)
		router.POST("/emailContact", controllers.Controller{}.EmailContact)
	}
}
