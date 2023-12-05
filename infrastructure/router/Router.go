package router

import (
	"go-project/adapters/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, c controller.AppController) {
	r.POST("/register", c.User.HandleRegister)
	auth := r.Group("/login")
	auth.Use(AuthMiddleWare())
	{
		auth.POST("", c.User.HandleRegister)
	}
}
