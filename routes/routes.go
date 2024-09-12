package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/talles-morais/gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.ShowOneStudent)
	r.GET("/students/:id", controllers.SearchStudentById)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.POST("/students", controllers.CreateStudent)
	r.Run()
}
