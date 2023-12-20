package main

import (
	"example/board/auth"
	"example/board/controllers"
	"example/board/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	loadErr := godotenv.Load(".env")

	if loadErr != nil {
		panic(loadErr)
	}

	secret := os.Getenv("JWT_SECRET")

	authMiddleware := auth.NewAuthentication(secret)

	openAPI := r.Group("/")
	{
		openAPI.POST("/signup", controllers.CreateUser)
		openAPI.POST("/login", controllers.Login)

		openAPI.GET("/posts/all", controllers.FindPosts)
		openAPI.GET("/posts/:id", controllers.FindPost)
	}

	secureAPI := r.Group("/")
	secureAPI.Use(authMiddleware.StripTokenMiddleware())
	{
		secureAPI.GET("/users", controllers.FindUser)
		secureAPI.PUT("/users", controllers.UpdateUser)
		secureAPI.DELETE("/users", controllers.DeleteUser)

		secureAPI.POST("/posts", controllers.CreatePost)
		secureAPI.PUT("/posts/:id", controllers.UpdatePost)
		secureAPI.DELETE("/posts/:id", controllers.DeletePost)
	}

	err := r.Run()
	if err != nil {
		return
	}
}
