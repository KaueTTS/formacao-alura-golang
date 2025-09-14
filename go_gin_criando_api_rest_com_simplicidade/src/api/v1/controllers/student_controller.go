package controllers_v1

import (
	postgresdb "go_gin_criando_api_rest_com_simplicidade/src/config/db"
	"go_gin_criando_api_rest_com_simplicidade/src/models"

	"github.com/gin-gonic/gin"
)

func DisplayAllStudents(c *gin.Context) {
	var students []models.Student
	postgresdb.DB.Find(&students)
	c.JSON(200, students)
}

func DisplayStudentByID(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := postgresdb.DB.First(&student, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(200, student)
}

func CreateStudent(c *gin.Context) {
	var newStudent models.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	postgresdb.DB.Create(&newStudent)
	c.JSON(201, newStudent)
}
