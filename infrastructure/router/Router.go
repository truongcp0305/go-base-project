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
		auth.POST("", c.User.HandleLogin)
	}
	r.GET("/cmd/list-process", c.Cmd.HandleListProcess)
	r.POST("/cmd/kill-process", c.Cmd.HandleKillProcess)
	r.POST("/cmd/execute-script", c.Cmd.HandleExecuteScript)
	r.POST("/cmd/open-file", c.Cmd.HandleOpenFile)
}
