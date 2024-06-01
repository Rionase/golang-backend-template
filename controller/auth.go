package controller

import (
	"golang-backend-template/model"
	"golang-backend-template/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Login(*gin.Context)
	RegisterUser(*gin.Context)
	RegisterAdmin(*gin.Context)
	Logout(*gin.Context)
}

type authController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *authController {
	return &authController{authService: authService}
}

func (a authController) Login(c *gin.Context) {
	var userData model.UserLoginBody
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	expiredAt := time.Now().Add(24 * time.Hour)
	token, err := a.authService.Login(userData, expiredAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.JwtResponse{
		Token:     token,
		ExpiredAt: expiredAt,
		Message:   "success login",
	})
}

func (a authController) RegisterUser(c *gin.Context) {
	var userData model.UserAuthBody
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	expiredAt := time.Now().Add(24 * time.Hour)
	token, err := a.authService.RegisterUser(userData, expiredAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.JwtResponse{
		Token:     token,
		ExpiredAt: expiredAt,
		Message:   "success register new user",
	})
}

func (a authController) RegisterAdmin(c *gin.Context) {
	var userData model.UserAuthBody
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	expiredAt := time.Now().Add(24 * time.Hour)
	token, err := a.authService.RegisterAdmin(userData, expiredAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.JwtResponse{
		Token:     token,
		ExpiredAt: expiredAt,
		Message:   "success register new admin",
	})
}

func (a authController) Logout(c *gin.Context) {
	var userData model.UserAuthBody
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if err := a.authService.Logout(userData); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "success logout",
	})
}
