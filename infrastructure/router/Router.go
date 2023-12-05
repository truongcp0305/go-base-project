package router

import (
	"go-project/adapters/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, c controller.AppController) {
	//r.Use(AuthMiddleWare())
	r.GET("/login", c.User.HandleLogin)
	r.POST("/register", c.User.HandleRegister)
}
