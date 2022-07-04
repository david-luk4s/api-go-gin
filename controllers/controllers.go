package controllers

import (
	"net/http"

	"github.com/david-luk4s/api-go-gin/database"
	"github.com/david-luk4s/api-go-gin/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Api Rest write in Gin Golang",
	})
}

// GetAllStudants godoc
// @Summary      List students
// @Description  get students
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   models.Student
// @Router       /students [get]
func GetAllStudants(c *gin.Context) {
	var student []models.Student
	database.DB.Find(&student)
	c.JSON(200, student)
}

// GetStudants godoc
// @Summary      Show an student
// @Description  get string by ID
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Student ID"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Router       /students/{id} [get]
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

// CreateStundent godoc
// @Summary      Add an students
// @Description  add by json students
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        students  body     models.Student  true  "Add Student"
// @Success      200      {object}  models.Student
// @Failure      400      {object}  httputil.HTTPError
// @Router       /students [post]
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

// DeleteStudent godoc
// @Summary      Delete an student
// @Description  Delete by student ID
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Student ID"  Format(int64)
// @Success      204  {object}  models.Student
// @Router       /students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	database.DB.Delete(&models.Student{}, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "student delete with sucess",
	})
}

// UpdateStudent godoc
// @Summary      Update an Student
// @Description  Update by json Student
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id       path      int                  true  "Student ID"
// @Param        student  body      models.Student  true  "Update Student"
// @Success      200      {object}  models.Student
// @Failure      400      {object}  httputil.HTTPError
// @Router       /students/{id} [patch]
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

// SearchByCPF godoc
// @Summary      List students
// @Description  get students
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        cpf  query     string  false  "name search by cpf"  Format(cpf)
// @Success      200  {array}   models.Student
// @Router       /students/cpf/:cpf	 [get]
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

// RenderPageNotFound godoc
// @Summary      Default page not found
// @Description  Default page for router not found
// @Success      400  {empty}   httputil.HTTPError
func RenderPageNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
