package main

import (
	"golang-backend-template/controller"
	"golang-backend-template/db"
	"golang-backend-template/model"
	"golang-backend-template/repository"
	"golang-backend-template/router"
	"golang-backend-template/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.NewPostgresDB()

	dbCredential := model.DatabaseCredential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "Saputra123#",
		DatabaseName: "testing",
		Port:         5432,
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.User{}, &model.Post{})

	userRepo := repository.NewUserRepo(conn)
	postRepo := repository.NewPostRepository(conn)

	authService := service.NewAuthService(userRepo)
	postService := service.NewPostService(postRepo)

	authController := controller.NewAuthController(authService)
	postController := controller.NewPostController(postService)

	route := gin.Default()

	router.AuthRoutes(route, authController)
	router.PostRoutes(route, postController)

	route.Run(":8080")
}
