package routers

import (
	"github.com/david-luk4s/api-go-gin/controllers"
	docs "github.com/david-luk4s/api-go-gin/docs"
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

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		students := v1.Group("/students")
		{
			students.GET("", controllers.GetAllStudants)
			students.GET(":id", controllers.GetStudants)
			students.POST("", controllers.CreateStundent)
			students.DELETE(":id", controllers.DeleteStudent)
			students.PATCH(":id", controllers.UpdateStudent)
			students.GET("/cpf/:cpf", controllers.SearchByCPF)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
