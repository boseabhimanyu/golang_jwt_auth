package httpserver

import (
	"golang-jwt-auth/internal/app"
	"golang-jwt-auth/internal/middleware"
	"golang-jwt-auth/internal/user"
	"net/http"

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

	//list all data/files (protected)
	api := r.Group("/api")

	//implementing middleware in routes

	api.Use(middleware.AuthRequired(a.Config.JWTSecret))

	api.GET("/files", func(c *gin.Context) {

		userID, _ := middleware.GetUserID(c)
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"userId": userID,
			"files":  []any{},
		})
	})

	api.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":       true,
			"products": []any{},
		})
	})

	admin := api.Group("/admin")

	admin.Use(middleware.RequireAdmin())

	admin.GET("/restricted", func(c *gin.Context) {
		role, _ := middleware.GetRole(c)
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"role": role,
		})
	})

	return r
}
