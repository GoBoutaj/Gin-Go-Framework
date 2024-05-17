package controllers

import (
	"github.com/gin-gonic/gin"
	"miboutaj.com/go-crud/initializers"
	"miboutaj.com/go-crud/models"
)

func StudentCreate(c *gin.Context) {
	//Get data from request body

	var body struct {
		Firstname string
		Lastname  string
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	//Create Student
	student := models.Student{
		FirstName: body.Firstname,
		LastName:  body.Lastname,
	}

	result := initializers.DB.Create(&student)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"student": student,
	})

}

func GetStudents(c *gin.Context) {
	var students []models.Student

	initializers.DB.Find(&students)

	c.JSON(200, gin.H{
		"students": students,
	})
}

func GetStudentById(c *gin.Context) {

	var student models.Student

	id := c.Param("id")

	initializers.DB.First(&student, id)

	c.JSON(200, gin.H{
		"student": student,
	})

}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Firstname string
		Lastname  string
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	var student models.Student
	initializers.DB.First(&student, id)
	initializers.DB.Model(&student).Updates(models.Student{
		FirstName: body.Firstname,
		LastName:  body.Lastname,
	})

	c.JSON(200, gin.H{
		"student": student,
	})

}
