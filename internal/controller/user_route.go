package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-jwt-exerices/internal/repository"
	"github.com/lipaysamart/go-jwt-exerices/internal/service"
	"github.com/lipaysamart/go-jwt-exerices/pkg/middleware"
	"github.com/lipaysamart/gocommon/dbs"
)

func UserRoute(r *gin.RouterGroup, db dbs.IDatabase) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandle := NewUserHandle(userService)

	authMiddleware := middleware.JWTAuth()
	refreshAuthMiddleware := middleware.JWTRefresh()
	userGroup := r.Group("/auth")
	{
		userGroup.POST("/register", userHandle.Register)
		userGroup.POST("/login", userHandle.Login)
		userGroup.POST("/me", authMiddleware, userHandle.GetMe)
		userGroup.POST("/refresh", refreshAuthMiddleware, userHandle.RefreshToken)
		// userGroup.POST("/profile/:id", userHandle.UpdateProfile)
	}
}
