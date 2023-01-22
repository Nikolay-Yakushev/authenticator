package httpapp

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func(a *Adapter) initRoutes(router *gin.Engine, logger *zap.Logger){
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	v1 := router.Group("/api/v1/")
	{
		v1.POST("/login", a.Login)
		v1.POST("/singup", a.Singup)
		v1.POST("/logout", a.Logout)
		v1.POST("/verify", a.Verify)
	}
}