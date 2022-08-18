package routes

import (
	"github.com/adel-habib/golang-jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes gin.Engine) {
	incomingRoutes.POST("user/signup", controllers.Signup())
	incomingRoutes.POST("user/signin", controllers.Signin())
}
