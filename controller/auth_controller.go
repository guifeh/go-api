package controller

import (
	"go-api/model"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userUseCase *usecase.UserUseCase
}

func NewAuthController(userUseCase *usecase.UserUseCase) *AuthController {
	return &AuthController{userUseCase: userUseCase}
}

func (a *AuthController) Register(ctx *gin.Context) {
	var request model.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.userUseCase.CreateUser(request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, user)
}

func (a *AuthController) Login(ctx *gin.Context) {
	var credentials model.Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := a.userUseCase.LoginUser(credentials)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}
