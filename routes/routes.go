package routes

import (
	"RESTAPICRUD/controllers"
	"RESTAPICRUD/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.Registar)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/profile", controllers.GetProfile)
	return r
}
