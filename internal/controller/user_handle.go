package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-jwt-exerices/internal/model"
	"github.com/lipaysamart/go-jwt-exerices/internal/service"
	"github.com/lipaysamart/go-jwt-exerices/pkg/utils"
)

type UserHandle struct {
	userHandle service.IUserService
}

func NewUserHandle(handle service.IUserService) *UserHandle {
	return &UserHandle{
		userHandle: handle,
	}
}

func (h *UserHandle) Login(ctx *gin.Context) {
	var httpReq model.UserLoginReq

	if err := ctx.ShouldBind(&httpReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed parse request body...",
			"error":   err.Error(),
		})
		return
	}

	user, accessToken, refreshToken, err := h.userHandle.Login(ctx, &httpReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to login user...",
			"error":   err.Error(),
		})
		return
	}

	var userLoginResp model.UserLoginResp
	userLoginResp.AccessToken = accessToken
	userLoginResp.RefreshToken = refreshToken
	utils.Copy(&userLoginResp.User, user)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login success...",
		"data":    userLoginResp,
	})
}
func (h *UserHandle) Register(ctx *gin.Context) {
	var httpReq model.UserRegisterReq

	if err := ctx.ShouldBind(&httpReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed parse request body...",
			"error":   err.Error(),
		})
		return
	}
	if err := h.userHandle.Register(ctx, &httpReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register user...",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register success...",
	})
}

func (h *UserHandle) GetMe(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to get user...",
			"error":   "Unauthorized",
		})
		return
	}

	user, err := h.userHandle.GetUserByID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get user...",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user success...",
		"data":    user,
	})
}

func (h *UserHandle) RefreshToken(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to get user...",
			"error":   "Unauthorized",
		})
		return
	}

	accessToken, err := h.userHandle.RefreshToken(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to refresh token...",
			"error":   err.Error(),
		})
		return
	}

	resp := model.RefreshTokenRes{
		AccessToken: accessToken,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Refresh token success...",
		"data":    resp,
	})

}
