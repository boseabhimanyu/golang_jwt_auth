package httpserver

import (
	"golang-jwt-auth/internal/app"
	"golang-jwt-auth/internal/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/health", health)

	userRepo := user.NewRepo(a.DB)
	userSvc := user.NewService(userRepo, a.Config.JWTSecret)

	userhandler := user.NewHandler(userSvc)
	// A -> unauth user
	// B -> auth + user (role)
	// C -> auth + admin (role)

	//unauth routes -> public routes

	r.POST("/register", userhandler.Register)
	r.POST("/login", userhandler.Login)

	return r
}
