package main

import (
	"log"

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
// @BasePath /api/v1

func main() {
	// Repository 초기화
	repo := fake.NewRepository()
	service := flow.NewService(repo)

	// Handler 초기화
	taskHandler := NewHandler(service)

	// Gin 라우터 설정
	r := gin.Default()

	// CORS 미들웨어 추가
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API 라우트 그룹 설정
	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("", taskHandler.ListTasks)
			tasks.GET("/:id", taskHandler.GetTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}
	}

	// Swagger 문서 라우트
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 헬스 체크 엔드포인트
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Tasker API is running",
		})
	})

	// 루트 경로에서 Swagger UI로 리다이렉트
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	log.Println("Starting Tasker API server on :8080")
	log.Println("Swagger UI available at: http://localhost:8080/swagger/index.html")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
