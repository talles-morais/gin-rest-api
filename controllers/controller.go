package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talles-morais/gin-rest-api/database"
	"github.com/talles-morais/gin-rest-api/models"
)

func SearchStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	c.JSON(200, students)
}

func ShowOneStudent(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API diz:": "E ai " + name + ", tudo beleza?",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	// error treatment
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso",
	})

}
