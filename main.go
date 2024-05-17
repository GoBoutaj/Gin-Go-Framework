package main

import (
	"github.com/gin-gonic/gin"
	"miboutaj.com/go-crud/controllers"
	"miboutaj.com/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/create-student", controllers.StudentCreate)
	r.GET("/students", controllers.GetStudents)
	r.GET("/student/:id", controllers.GetStudentById)
	r.PUT("/student/update/:id", controllers.UpdateStudent)
	r.Run()
}
