package controllers

import (
	"net/http"

	"go_api_rest/database"
	"go_api_rest/models"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

func Saudacoes(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", Tudo beleza?",
	})
}

// TodosAlunos godoc
//
//	@Summary		Exibir todos os alunos
//	@Description	Rota para exibir todos os alunos
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Aluno
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/alunos	[get]
func TodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func CriarNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscarAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

// EditarAlunos godoc
//
//	@Summary		Edita um aluno
//	@Description	Rota para editar um aluno
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Param			alunos	body	models.Aluno	true	"Modelo do aluno"
//	@Success		200	{object}	models.Aluno
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/alunos/{id}	[patch]
func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	database.DB.Save(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
