package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type HandlerHTTP struct {
	uuc domain.UserUsecase
	suc domain.SessionUsecase
}

func NewHandlerHTTP(
	router *gin.Engine,
	uuc domain.UserUsecase,
	suc domain.SessionUsecase,
	tokenChecker gin.HandlerFunc,
) *HandlerHTTP {
	handler := &HandlerHTTP{
		uuc: uuc,
		suc: suc,
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", handler.signIn)
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/refresh", handler.refresh)
		auth.POST("/logout", tokenChecker, handler.logout)
	}

	return handler
}

func generateFingerprint(c *gin.Context) string {
	return fmt.Sprintf(
		"%s; %s",
		c.Request.Header["User-Agent"],
		c.Request.Header["Accept-Language"],
	)
}
