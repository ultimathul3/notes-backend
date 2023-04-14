package notebook

import (
	"net/http"
	"strconv"

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
		notebook.GET("/", handler.getAllByUserID)
		notebook.PUT("/:id", handler.update)
		notebook.DELETE("/:id", handler.delete)
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

func (h *HandlerHTTP) getAllByUserID(c *gin.Context) {
	userID := c.MustGet("userID").(int64)

	notebooks, err := h.nuc.GetAllByUserID(c, userID)
	if err != nil {
		log.Error("get all notebooks by user id: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notebooks)
}

func (h *HandlerHTTP) update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid id param"})
		return
	}

	var notebook domain.Notebook
	if err := c.BindJSON(&notebook); err != nil {
		log.Error("Notebook bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)
	notebook.ID = id
	notebook.UserID = userID

	err = h.nuc.Update(c, notebook)
	if err != nil {
		log.Error("update notebook: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *HandlerHTTP) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid id param"})
		return
	}

	userID := c.MustGet("userID").(int64)
	err = h.nuc.Delete(c, id, userID)
	if err != nil {
		log.Error("delete notebook: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "notebook not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
