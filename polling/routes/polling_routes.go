package routes

import (
	"polling/polling/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	
	route := router.Group("/polling")
	{
		route.GET("/status", controllers.CheckStatusInformationStudent)
		route.GET("/students", controllers.CheckListStudents)
	}

}