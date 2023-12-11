package controller

import (
	"errors"
	"go-project/model"
	"go-project/usecase/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) UserController {
	return UserController{
		userService: us,
	}
}

// @Summary Login
// @Description Login to app
// @Tags User
// @Security BearerAuth
// @Param Body body model.User true "query params"
// @Success 200 {object} outgoing.LoginOutgoing
// @Failure 400 {object} outgoing.ModelBadRequestErr
// @Failure 500 {object} outgoing.ModelInternalErr
// @Router /login [post]
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
	err := u.userService.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

// @Summary Register
// @Description Register new account
// @Tags User
// @Param Body body model.User true "query params"
// @Success 200 {object} outgoing.RegisterOutgoing
// @Failure 400 {object} outgoing.ModelBadRequestErr
// @Failure 500 {object} outgoing.ModelInternalErr
// @Router /register [post]
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
	err := u.userService.CreateAccount(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
