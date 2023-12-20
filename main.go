package main

import (
	"example/board/auth"
	"example/board/controllers"
	"example/board/models"
	"example/board/models/repositories"
	"example/board/services"
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

	userController := controllers.UserController{
		UserService: services.UserService{
			UserRepository: repositories.NewUserRepository(models.DB),
			PostRepository: repositories.NewPostRepository(models.DB),
		},
	}

	postController := controllers.PostController{
		PostService: services.PostService{
			PostRepository: repositories.NewPostRepository(models.DB),
		},
	}

	openAPI := r.Group("/")
	{
		openAPI.POST("/signup", userController.CreateUser)
		openAPI.POST("/login", controllers.Login)

		openAPI.GET("/posts/all", postController.FindPosts)
		openAPI.GET("/posts/:id", postController.FindPost)
	}

	secureAPI := r.Group("/")
	secureAPI.Use(authMiddleware.StripTokenMiddleware())
	{
		secureAPI.GET("/users", userController.FindUser)
		secureAPI.PUT("/users", userController.UpdateUser)
		secureAPI.DELETE("/users", userController.DeleteUser)

		secureAPI.POST("/posts", postController.CreatePost)
		secureAPI.PUT("/posts/:id", postController.UpdatePost)
		secureAPI.DELETE("/posts/:id", postController.DeletePost)
	}

	err := r.Run()
	if err != nil {
		return
	}
}
