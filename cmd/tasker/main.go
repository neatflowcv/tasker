package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neatflowcv/tasker/internal/app/flow"
	"github.com/neatflowcv/tasker/internal/pkg/repository/fake"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/neatflowcv/tasker/docs" // Import for side effects
)

// @title Tasker API
// @version 1.0
// @description Task 관리를 위한 REST API
// @host localhost:8080
// @BasePath /tasker/v1

func main() {
	// Repository 초기화
	repo := fake.NewRepository()
	service := flow.NewService(repo)

	// Handler 초기화
	taskHandler := NewHandler(service)

	// Gin 라우터 설정
	router := gin.Default()

	// CORS 미들웨어 추가
	router.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)

			return
		}

		ctx.Next()
	})

	// API 라우트 그룹 설정
	v1 := router.Group("/tasker/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("", taskHandler.ListTasks)
			tasks.GET("/:id", taskHandler.GetTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}

		// Swagger 문서 라우트
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/tasker/v1/swagger/index.html")
		})
	}

	log.Println("Starting Tasker API server on :8080")
	log.Println("Swagger UI available at: http://localhost:8080/swagger/index.html")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
