package note

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type HandlerHTTP struct {
	nuc domain.NoteUsecase
}

func NewHandlerHTTP(router *gin.Engine, nuc domain.NoteUsecase, tokenChecker gin.HandlerFunc) {
	handler := &HandlerHTTP{
		nuc: nuc,
	}

	notebook := router.Group("/notebooks/:id/notes").Use(tokenChecker)
	{
		notebook.POST("/", handler.create)
		notebook.GET("/", handler.getAllByNotebookID)
	}
}

// @Summary		Creating a note in notebook
// @Security	BearerToken
// @Tags		Note
// @Accept		json
// @Produce		json
// @Param		notebook_id path int true "Notebook ID"
// @Param		user body domain.CreateNoteDTO true "Note data"
// @Success		200 {object} docs.CreateNoteResponse "Note ID"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/notebooks/{notebook_id}/notes [post]
func (h *HandlerHTTP) create(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	var note domain.CreateNoteDTO
	if err := c.BindJSON(&note); err != nil {
		log.Error("Note bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)

	id, err := h.nuc.Create(c, userID, notebookID, note)
	if err != nil {
		log.Error("create note: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary		Getting a list of user notes in notebook
// @Security	BearerToken
// @Tags		Note
// @Accept		json
// @Produce		json
// @Param		notebook_id path int true "Notebook ID"
// @Success		200 {array} domain.GetAllNotesResponse "Notebooks"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/notebooks/{notebook_id}/notes [get]
func (h *HandlerHTTP) getAllByNotebookID(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	notes, err := h.nuc.GetAllByNotebookID(c, userID, notebookID)
	if err != nil {
		log.Error("get all notes by notebook id: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.GetAllNotesResponse{
		Notes: notes,
		Count: len(notes),
	})
}
