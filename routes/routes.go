package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/talles-morais/gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.ShowOneStudent)
	r.GET("/students/:id", controllers.SearchStudentById)
	r.GET("/students/cpf/:cpf", controllers.SearchStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.POST("/students", controllers.CreateStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/home", controllers.RenderIndexPage)
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}
