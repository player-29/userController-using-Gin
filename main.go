package main

import (
	db "go-sqlite-crud-project/config"
	"go-sqlite-crud-project/controller"
	"go-sqlite-crud-project/repository"
	"go-sqlite-crud-project/service"
    "github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db.InitializeDatabase()

	// Create repository, service, and controller
	userRepo := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.GetUser)
	r.GET("/users", userController.GetAllUsers)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	// Start server
	r.Run(":8080")

}