package notebook

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type HandlerHTTP struct {
	nuc domain.NotebookUsecase
}

func NewHandlerHTTP(router *gin.Engine, nuc domain.NotebookUsecase, tokenChecker gin.HandlerFunc) {
	handler := &HandlerHTTP{
		nuc: nuc,
	}

	notebook := router.Group("/notebooks").Use(tokenChecker)
	{
		notebook.POST("/", handler.create)
	}
}

func (h *HandlerHTTP) create(c *gin.Context) {
	var notebook domain.Notebook
	if err := c.BindJSON(&notebook); err != nil {
		log.Error("Notebook bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)
	notebook.UserID = userID

	id, err := h.nuc.Create(c, notebook)
	if err != nil {
		log.Error("create notebook: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
