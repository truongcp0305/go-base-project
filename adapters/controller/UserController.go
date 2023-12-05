package controller

import (
	"errors"
	"go-project/model"
	"go-project/usecase/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	loginService    service.LoginService
	registerService service.RegisterService
}

func NewUserController(l service.LoginService, r service.RegisterService) UserController {
	return UserController{
		loginService:    l,
		registerService: r,
	}
}

func (u *UserController) HandleLogin(c *gin.Context) {
	user := model.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.UserName == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Invalid params")})
		return
	}
	err := u.loginService.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (u *UserController) HandleRegister(c *gin.Context) {
	user := model.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.DisplayName == "" || user.UserName == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Invalid params")})
		return
	}
	err := u.registerService.CreateAccount(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
