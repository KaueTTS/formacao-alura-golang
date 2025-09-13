package route

import (
	controllers_v1 "go_gin_criando_api_rest_com_simplicidade/src/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	group := r.Group("/v1/students")
	{
		group.GET("/", controllers_v1.DisplayAllStudents)
		group.POST("/", controllers_v1.CreateStudent)
	}
}
