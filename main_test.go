package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/talles-morais/gin-rest-api/controllers"
	"github.com/talles-morais/gin-rest-api/database"
	"github.com/talles-morais/gin-rest-api/models"
)

var ID int

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router
}

func CreateMockStudent() {
	student := models.Student{
		Name: "Fernando",
		CPF: "08008008012",
		Phone: "24998611506",
	}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteMockStudent() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestGetOneStudent(t *testing.T) {
	r := SetupRouter()
	r.GET("/:name", controllers.ShowOneStudent)
	request, _ := http.NewRequest("GET", "/talles", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")
	responseMock := `{"API diz:":"E ai talles, tudo beleza?"}`
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, responseMock, string(responseBody))
}

func TestGetAllStudents(t *testing.T) {
	database.Connect()
	CreateMockStudent()
	defer DeleteMockStudent()
	r := SetupRouter()
	r.GET("/students", controllers.ShowAllStudents)

	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchStudent(t *testing.T) {
	database.Connect()
	CreateMockStudent()
	defer DeleteMockStudent()
	r := SetupRouter()
	r.GET("/students/cpf/:cpf", controllers.SearchStudent)
	request, _ := http.NewRequest("GET", "/students/cpf/08008008012", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchStudentById(t *testing.T) {
	database.Connect()
	CreateMockStudent()
	defer DeleteMockStudent()

	r:=SetupRouter()
	r.GET("/students/:id", controllers.SearchStudentById)
	path := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var mockStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &mockStudent)
	assert.Equal(t, "Fernando", mockStudent.Name)
	assert.Equal(t, "08008008012", mockStudent.CPF)
	assert.Equal(t, "24998611506", mockStudent.Phone)
}

func TestDeleteStudent(t *testing.T) {
	database.Connect()
	CreateMockStudent()
	r := SetupRouter()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	path := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}