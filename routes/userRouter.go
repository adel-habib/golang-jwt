package routes

import (
	"github.com/adel-habib/golang-jwt/controllers"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
