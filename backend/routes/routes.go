package routes

import (
	"jopanel/backend/controllers"
	"jopanel/backend/services"
	"jopanel/backend/middleware"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Services
	jwtService := services.NewJWTService()
	userService := services.NewUserService()
	
	authController := controllers.NewAuthController(jwtService)
	userController := controllers.NewUserController(userService)
	fileController := controllers.NewFileController(services.NewFileService())

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
		
		// Auth Routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)
		}

		// Admin Routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.RequireRole("admin"))
		{
			admin.POST("/users", userController.CreateUser)
			admin.GET("/users", userController.GetAllUsers)
			admin.GET("/users/:id", userController.GetUser)
			admin.POST("/users/:id/suspend", userController.SuspendUser)
			admin.POST("/users/:id/unsuspend", userController.UnsuspendUser)
		}

		// User Routes
		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware())
		{
			user.GET("/files/list", fileController.ListFiles)
			user.GET("/files/content", fileController.GetContent)
			user.POST("/files/upload", fileController.Upload)
			user.POST("/files/mkdir", fileController.Mkdir)
			user.DELETE("/files/delete", fileController.Delete)
		}
	}
}
