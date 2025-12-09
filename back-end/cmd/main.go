package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/blog/back-end/db"
	"github.com/rillmind/blog/back-end/posts"
)

func main() {
	server := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "2306"
	}

	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	postsService := posts.NewService(dbConnection)
	postsController := posts.NewController(postsService)

	posts.RegisterRoutes(server, &postsController)

	// userRepository := user.NewRepository(dbConnection)
	// userService := user.NewService(userRepository)
	// userController := user.NewController(userService)

	// user.RegisterRoutes(server, &userController)

	server.Run(":" + port)
}
