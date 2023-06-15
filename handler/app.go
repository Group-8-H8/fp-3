package handler

import (
	"log"
	"os"

	"github.com/Group-8-H8/fp-3/database"
	_ "github.com/Group-8-H8/fp-3/docs"
	"github.com/Group-8-H8/fp-3/handler/http_handler"
	"github.com/Group-8-H8/fp-3/repository/category_repository/category_pg"
	"github.com/Group-8-H8/fp-3/repository/task_repository/task_pg"
	"github.com/Group-8-H8/fp-3/repository/user_repository/user_pg"
	"github.com/Group-8-H8/fp-3/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

var PORT string

// @title           Final Project 3 - Group 8 Hacktiv8
// @version         1.0
// @description     This is a documentation for kanban board API from final project 3 - Group 8 Hacktiv8

// @host      localhost:8080
// @BasePath  /api/v1

func StartApp() {
	db := database.GetPostgresInstance()

	r := gin.Default()
	route := r.Group("/api/v1")

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := http_handler.NewUserHandler(userService)
	authService := service.NewAuthService(userRepo)

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)

		userRoute.Use(authService.Authentication())
		userRoute.PUT("/update-account", userHandler.UpdateAccount)
		userRoute.DELETE("/delete-account", userHandler.DeleteAccount)
	}

	categoryRepo := category_pg.NewCategoryPg(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := http_handler.NewCategoryHandler(categoryService)

	categoryRoute := route.Group("/categories")
	{
		categoryRoute.Use(authService.Authentication())
		categoryRoute.POST("/", authService.Authorization(), categoryHandler.CreateCategory)
		categoryRoute.PATCH("/:categoryId", authService.Authorization(), categoryHandler.UpdateCategory)
		categoryRoute.GET("/", categoryHandler.GetCategories)
		categoryRoute.GET("/:categoryId", categoryHandler.GetCategory)
		categoryRoute.DELETE("/:categoryId", authService.Authorization(), categoryHandler.DeleteCategory)
	}

	taskRepo := task_pg.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo, categoryRepo, userRepo)
	taskHandler := http_handler.NewTaskHandler(taskService)

	taskRoute := route.Group("/tasks")
	{
		taskRoute.Use(authService.Authentication())
		taskRoute.POST("/", taskHandler.CreateTask)
		taskRoute.GET("/", taskHandler.GetTasks)
		taskRoute.GET("/:taskId", taskHandler.GetTask)
		taskRoute.PUT("/:taskId", taskHandler.UpdateTask)
		taskRoute.PATCH("/update-status/:taskId", taskHandler.UpdateTasksStatus)
		taskRoute.PATCH("/update-category/:taskId", taskHandler.UpdateTasksCategory)
		taskRoute.DELETE("/:taskId", taskHandler.DeleteTask)
	}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3000"
	}
	log.Fatalln(r.Run(":" + PORT))
}
