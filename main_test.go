package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/david-luk4s/api-go-gin/controllers"
	"github.com/david-luk4s/api-go-gin/database"
	"github.com/david-luk4s/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func EngineRouterTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routers := gin.Default()
	return routers
}

func CreateStudentTest() {
	student := models.Student{Name: "Test", RG: "123456789", CPF: "03159984230"}
	database.ConnectionDBTest()
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentTest() {
	database.ConnectionDBTest()
	database.DB.Delete(&models.Student{}, ID)
}

func TestListStudents(t *testing.T) {
	database.ConnectionDBTest()
	CreateStudentTest()
	defer DeleteStudentTest()

	r := EngineRouterTest()
	r.GET("/students", controllers.GetAllStudants)

	req, _ := http.NewRequest("GET", "/students", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestSearchCPF(t *testing.T) {
	database.ConnectionDBTest()
	CreateStudentTest()
	defer DeleteStudentTest()

	r := EngineRouterTest()
	r.GET("/students/cpf/:cpf", controllers.SearchByCPF)

	req, _ := http.NewRequest("GET", "/students/cpf/03159984230", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetStudents(t *testing.T) {
	CreateStudentTest()
	defer DeleteStudentTest()
	r := EngineRouterTest()
	r.GET("/students/:id", controllers.GetStudants)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var student models.Student
	json.Unmarshal(resp.Body.Bytes(), &student)

	assert.Equal(t, ID, int(student.ID))
	assert.Equal(t, "Test", student.Name)
	assert.Equal(t, "123456789", student.RG)
	assert.Equal(t, "03159984230", student.CPF)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectionDBTest()
	CreateStudentTest()
	r := EngineRouterTest()
	r.DELETE("/stundents/:id", controllers.DeleteStudent)
	path := "/stundents/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateStudent(t *testing.T) {
	database.ConnectionDB()
	CreateStudentTest()
	defer DeleteStudentTest()
	r := EngineRouterTest()
	r.PATCH("/students/:id", controllers.UpdateStudent)

	path := "/students/" + strconv.Itoa(ID)
	student := models.Student{Name: "Test", RG: "123456700", CPF: "47159984230"}
	alunojson, _ := json.Marshal(&student)

	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(alunojson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var studentupdate models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentupdate)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, student.Name, studentupdate.Name)
	assert.Equal(t, student.CPF, studentupdate.CPF)
	assert.Equal(t, student.RG, studentupdate.RG)
}
