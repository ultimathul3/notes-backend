package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type HandlerHTTP struct {
	uuc domain.UserUsecase
}

func NewHandlerHTTP(router *gin.Engine, uuc domain.UserUsecase) {
	handler := &HandlerHTTP{
		uuc: uuc,
	}

	users := router.Group("/users")
	{
		users.POST("/", handler.create)
		users.GET("/:id", handler.getByID)
	}
}

func (h *HandlerHTTP) create(c *gin.Context) {

}

func (h *HandlerHTTP) getByID(c *gin.Context) {

}
