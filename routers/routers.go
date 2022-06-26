package routers

import (
	"github.com/david-luk4s/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("", controllers.Index)
	r.NoRoute(controllers.RenderPageNotFound)
	r.GET("/students", controllers.GetAllStudants)
	r.GET("/students/:id", controllers.GetStudants)
	r.POST("/students", controllers.CreateStundent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchByCPF)
	r.Run()
}
