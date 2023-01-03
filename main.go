package main

import (
	"diary_api/controller"
	"diary_api/database"
	"diary_api/middleware"
	"diary_api/schema"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&schema.Role{})
	database.Database.AutoMigrate(&schema.User{})
	database.Database.AutoMigrate(&schema.UserRoles{})
	// database.Database.AutoMigrate(&schema.Entries{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)
	protectedRoutes.GET("/entry/:ID", controller.FindEntry)
	protectedRoutes.PATCH("/entry/:ID/edit", controller.UpdateContent)

	// Role //
	specialRoute := router.Group("/special")
	specialRoute.POST("/add_role", controller.AddRole)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
