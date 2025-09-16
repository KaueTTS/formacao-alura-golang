package route

import (
	controllers "go_gin_criando_api_rest_com_simplicidade/src/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	group := r.Group("/v1/students")
	{
		group.GET("", controllers.DisplayAllStudents)
		group.GET("/:id", controllers.DisplayStudentByID)
		group.GET("/cpf/:cpf", controllers.DisplayStudentByCPF)

		group.POST("", controllers.CreateStudent)

		group.DELETE("/:id", controllers.DeleteStudent)

		group.PATCH("/:id", controllers.UpdateStudent)
	}
}
