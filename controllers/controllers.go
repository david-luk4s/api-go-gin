package controllers

import (
	"net/http"

	"github.com/david-luk4s/api-go-gin/database"
	"github.com/david-luk4s/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Api Rest write in Gin Golang",
	})
}

func RenderPageNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

func GetAllStudants(c *gin.Context) {
	var student []models.Student
	database.DB.Find(&student)
	c.JSON(200, student)
}

func GetStudants(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student

	database.DB.First(&student, id)

	if student.Model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Not found stundent",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStundent(c *gin.Context) {
	var student models.Student

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
	id := c.Params.ByName("id")
	database.DB.Delete(&models.Student{}, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "student delete with sucess",
	})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func SearchByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}
