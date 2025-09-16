package controllers_v1

import (
	postgresdb "go_gin_criando_api_rest_com_simplicidade/src/config/db"
	"go_gin_criando_api_rest_com_simplicidade/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayAllStudents(c *gin.Context) {
	var students []models.Student
	postgresdb.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func DisplayStudentByID(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := postgresdb.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func DisplayStudentByCPF(c *gin.Context) {
	cpf := c.Param("cpf")
	var student models.Student
	if err := postgresdb.DB.Where("cpf = ?", cpf).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var newStudent models.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postgresdb.DB.Create(&newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := postgresdb.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	postgresdb.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	postgresdb.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postgresdb.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}
