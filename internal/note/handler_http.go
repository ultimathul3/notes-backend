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

func NewHandlerHTTP(router *gin.Engine, nuc domain.NoteUsecase, tokenChecker gin.HandlerFunc) *HandlerHTTP {
	handler := &HandlerHTTP{
		nuc: nuc,
	}

	note := router.Group("/notebooks/:notebook-id/notes").Use(tokenChecker)
	{
		note.POST("/", handler.create)
		note.GET("/", handler.getAllByNotebookID)
		note.GET("/:note-id", handler.getByID)
		note.PATCH("/:note-id", handler.patch)
		note.DELETE("/:note-id", handler.delete)
	}

	return handler
}

// @Summary		Creating a note in a notebook
// @Security	BearerToken
// @Tags		Note
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		user body domain.CreateNoteDTO true "Note data"
// @Success		200 {object} docs.CreateNoteResponse "Note ID"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id}/notes/ [post]
func (h *HandlerHTTP) create(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	var note domain.CreateNoteDTO
	if err := c.BindJSON(&note); err != nil {
		log.Error("CreateNoteDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)

	id, err := h.nuc.Create(c, userID, notebookID, note)
	if err != nil {
		log.Error("create note: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrNotebookNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary		Getting a list of user notes in a notebook
// @Security	BearerToken
// @Tags		Note
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Success		200 {array} docs.GetAllNotesResponse "Notes"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id}/notes/ [get]
func (h *HandlerHTTP) getAllByNotebookID(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	notes, err := h.nuc.GetAllByNotebookID(c, userID, notebookID)
	if err != nil {
		log.Error("get all notes by notebook id: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrNotebookNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.GetAllNotesResponse{
		Notes: notes,
		Count: len(notes),
	})
}

// @Summary		Getting a note in a notebook
// @Security	BearerToken
// @Tags		Note
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		note-id path int true "Note ID"
// @Success		200 {array} docs.GetNoteResponse "Notes"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id}/notes/{note-id} [get]
func (h *HandlerHTTP) getByID(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	noteID, err := strconv.ParseInt(c.Param("note-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid note id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	note, err := h.nuc.GetByID(c, userID, notebookID, noteID)
	if err != nil {
		log.Error("get note by id: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrNoteNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}

// @Summary		Updating a note in a notebook
// @Security	BearerToken
// @Tags		Note
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		note-id path int true "Note ID"
// @Param		user body domain.PatchNoteDTO true "New note data"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id}/notes/{note-id} [patch]
func (h *HandlerHTTP) patch(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	noteID, err := strconv.ParseInt(c.Param("note-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid note id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	var note domain.PatchNoteDTO
	if err := c.BindJSON(&note); err != nil {
		log.Error("PatchNoteDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = h.nuc.Patch(c, noteID, userID, notebookID, note)
	if err != nil {
		log.Error("patch note: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @Summary		Deleting a note from a notebook
// @Security	BearerToken
// @Tags		Note
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		note-id path int true "Note ID"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id}/notes/{note-id} [delete]
func (h *HandlerHTTP) delete(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	noteID, err := strconv.ParseInt(c.Param("note-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid note id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	err = h.nuc.Delete(c, noteID, userID, notebookID)
	if err != nil {
		log.Error("delete note: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrNoteNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
