package routers

import (
	"github.com/david-luk4s/api-go-gin/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	r.Use(cors.Default())
	r.LoadHTMLGlob("templates/*")
	r.NoRoute(controllers.RenderPageNotFound)

	v1 := r.Group("/api/v1")
	{
		students := v1.Group("/students")
		{
			students.GET("/students", controllers.GetAllStudants)
			students.GET("/students/:id", controllers.GetStudants)
			students.POST("/students", controllers.CreateStundent)
			students.DELETE("/students/:id", controllers.DeleteStudent)
			students.PATCH("/students/:id", controllers.UpdateStudent)
			students.GET("/students/cpf/:cpf", controllers.SearchByCPF)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
