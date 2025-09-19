package routes

import (
	"go_api_rest/controllers"

	"github.com/gin-gonic/gin"

	docs "go_api_rest/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequest() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:nome", controllers.Saudacoes)
	r.GET("/alunos", controllers.TodosAlunos)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorID)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/alunos/", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
