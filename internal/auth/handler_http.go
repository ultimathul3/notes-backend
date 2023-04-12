package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type HandlerHTTP struct {
	uuc domain.UserUsecase
	suc domain.SessionUsecase
}

func NewHandlerHTTP(router *gin.Engine, uuc domain.UserUsecase, suc domain.SessionUsecase) {
	handler := &HandlerHTTP{
		uuc: uuc,
		suc: suc,
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", handler.signIn)
		auth.POST("/sign-up", handler.signUp)
	}
}
