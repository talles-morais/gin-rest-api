package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"github.com/talles-morais/gin-rest-api/database"
	"github.com/talles-morais/gin-rest-api/models"
)

// SearchStudentById godoc
// @Summary					 Procura aluno por id
// @Description			 Rota para procurar um aluno por sua id
// @Tags						 Students
// @Accept					 json
// @Produce					 json
// @Param   				 id	 	path 			int	true	"Id de aluno"
// @Success					 200  {object} 	models.Student
// @Failure					 400	{object}	httputil.HTTPError
// @Router					 /students/{id}	[get]
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

// ShowAllStudents 	 godoc
// @Summary					 Mostra todos os alunos
// @Description			 Rota para mostrar todos os alunos
// @Tags						 Students
// @Accept					 json
// @Produce					 json
// @Success					 200  {object}  models.Student
// @Failure					 400	{object}	httputil.HTTPError
// @Router					 /students	[get]
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

// CreateStudent 		 godoc
// @Summary					 Registra novo aluno
// @Description			 Rota para adicionar um aluno
// @Tags						 Students
// @Accept					 json
// @Produce					 json
// @Param   				 student	body	models.Student true "Modelo de aluno"
// @Success					 200 	{object} models.Student
// @Failure					 400	{object}	httputil.HTTPError
// @Router					 /students	[post]
func CreateStudent(c *gin.Context) {
	var student models.Student

	// error treatment
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

// SearchStudentById godoc
// @Summary					 deleta aluno por id
// @Description			 Rota para deletar um aluno por sua id
// @Tags						 Students
// @Accept					 json
// @Produce					 json
// @Param   				 id	 	path 			int	true	"Id de aluno"
// @Success					 200  {object} 	models.Student
// @Failure					 400	{object}	httputil.HTTPError
// @Router					 /students/{id}	[delete]
func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso",
	})

}

// SearchStudentById godoc
// @Summary					 edita aluno por id
// @Description			 Rota para editar um aluno por sua id
// @Tags						 Students
// @Accept					 json
// @Produce					 json
// @Param   				 id	 	path 			int	true	"ID do aluno"
// @Param   				 student	body	models.Student true "Modelo de aluno"
// @Success					 200  {object} 	models.Student
// @Failure					 400	{object}	httputil.HTTPError
// @Router					 /students/{id}	[patch]
func EditStudent(c *gin.Context) {
	var student models.Student
	id :=	c.Params.ByName("id")

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Model(&student).Where("id = ?", id).Updates(student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update student",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}


// SearchStudent godoc
// @Summary					 Procura aluno por CPF
// @Description			 Rota para procurar um aluno pelo CPF
// @Tags						 Students
// @Accept					 json
// @Produce					 json
// @Param   				 cpf	 	path 			string	true	"CPF do aluno"
// @Success					 200  {object}  models.Student
// @Failure					 404	{object}	httputil.HTTPError
// @Router					 /students/cpf/{cpf}	[get]
func SearchStudent(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func RenderIndexPage(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": students,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}