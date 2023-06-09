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

func NewHandlerHTTP(router *gin.Engine, nuc domain.NotebookUsecase, tokenChecker gin.HandlerFunc) *HandlerHTTP {
	handler := &HandlerHTTP{
		nuc: nuc,
	}

	notebook := router.Group("/notebooks").Use(tokenChecker)
	{
		notebook.POST("/", handler.create)
		notebook.GET("/", handler.getAllByUserID)
		notebook.PUT("/:notebook-id", handler.update)
		notebook.DELETE("/:notebook-id", handler.delete)
	}

	return handler
}

// @Summary		Creating notebook
// @Security	BearerToken
// @Tags		Notebook
// @Accept		json
// @Produce		json
// @Param		user body docs.CreateUpdateNotebookDTO true "Notebook data"
// @Success		200 {object} docs.CreateNotebookResponse "Notebook ID"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/ [post]
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

// @Summary		Getting a list of user notebooks
// @Security	BearerToken
// @Tags		Notebook
// @Accept		json
// @Produce		json
// @Success		200 {array} docs.GetAllNotebooksResponse "Notebooks"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/ [get]
func (h *HandlerHTTP) getAllByUserID(c *gin.Context) {
	userID := c.MustGet("userID").(int64)

	notebooks, err := h.nuc.GetAllByUserID(c, userID)
	if err != nil {
		log.Error("get all notebooks by user id: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.GetAllNotebooksResponse{
		Notebooks: notebooks,
		Count:     len(notebooks),
	})
}

// @Summary		Updating user notebook
// @Security	BearerToken
// @Tags		Notebook
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		user body docs.CreateUpdateNotebookDTO true "New notebook data"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id} [put]
func (h *HandlerHTTP) update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrNotebookNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @Summary		Deleting user notebook
// @Security	BearerToken
// @Tags		Notebook
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id} [delete]
func (h *HandlerHTTP) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	userID := c.MustGet("userID").(int64)
	err = h.nuc.Delete(c, id, userID)
	if err != nil {
		log.Error("delete notebook: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrNotebookNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
