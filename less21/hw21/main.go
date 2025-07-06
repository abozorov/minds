package main

import (
	"hw21/handler"

	"github.com/gin-gonic/gin"
	_ "hw21/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User CRUD API
// @version 1.0
// @description API для управления пользователями
// @host localhost:8000
// @BasePath
func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CRUD endpoints
	r.POST("/users", handler.CreateUser)
	r.GET("/users", handler.GetUsers)
	r.GET("/users/:id", handler.GetUserByID)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)
	r.Run(":8000")
	// http://localhost:8000/swagger/index.html
}
