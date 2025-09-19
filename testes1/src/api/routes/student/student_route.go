package route

import (
	controllers_v1 "testes1/src/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	group := r.Group("/v1/students")
	{
		group.GET("", controllers_v1.DisplayAllStudents)
		group.GET("/:id", controllers_v1.DisplayStudentByID)
		group.GET("/cpf/:cpf", controllers_v1.DisplayStudentByCPF)

		group.POST("", controllers_v1.CreateStudent)

		group.DELETE("/:id", controllers_v1.DeleteStudent)

		group.PATCH("/:id", controllers_v1.UpdateStudent)
	}
}
