package sharednote

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type HandlerHTTP struct {
	suc domain.SharedNoteUsecase
	uuc domain.UserUsecase
}

func NewHandlerHTTP(
	router *gin.Engine,
	suc domain.SharedNoteUsecase,
	uuc domain.UserUsecase,
	tokenChecker gin.HandlerFunc,
) *HandlerHTTP {

	handler := &HandlerHTTP{
		suc: suc,
		uuc: uuc,
	}

	sharedNote := router.Group("/shared-notes").Use(tokenChecker)
	{
		sharedNote.POST("/incoming", handler.create)
		sharedNote.GET("/incoming", handler.getIncomingSharedNotes)
		sharedNote.DELETE("/incoming/:shared-note-id", handler.delete)
	}

	return handler
}

// @Summary		Creating a shared note
// @Security	BearerToken
// @Tags		Shared note
// @Accept		json
// @Produce		json
// @Param		user body domain.CreateSharedNoteDTO true "Shared note data"
// @Success		200 {object} docs.CreateSharedNoteResponse "Shared note ID"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/shared-notes/incoming [post]
func (h *HandlerHTTP) create(c *gin.Context) {
	var sharedNote domain.CreateSharedNoteDTO
	if err := c.BindJSON(&sharedNote); err != nil {
		log.Error("CreateSharedNoteDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := sharedNote.Validate(); err != nil {
		log.Error("CreateSharedNoteDTO validate: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	whomID, err := h.uuc.GetUserIdByLogin(c, *sharedNote.Login)
	if err != nil {
		log.Error("get user id by login: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrUserNotFound.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)

	id, err := h.suc.Create(c, userID, whomID, *sharedNote.NoteID)
	if err != nil {
		log.Error("create shared note: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary		Deleting a shared note
// @Security	BearerToken
// @Tags		Shared note
// @Accept		json
// @Produce		json
// @Param		shared-note-id path int true "Shared note ID"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/shared-notes/incoming/{shared-note-id} [delete]
func (h *HandlerHTTP) delete(c *gin.Context) {
	sharedNoteID, err := strconv.ParseInt(c.Param("shared-note-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	err = h.suc.Delete(c, sharedNoteID, userID)
	if err != nil {
		log.Error("delete shared note: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrSharedNoteNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @Summary		Getting a list of incoming shared notes
// @Security	BearerToken
// @Tags		Shared note
// @Accept		json
// @Produce		json
// @Success		200 {array} domain.GetAllIncomingSharedNotesResponse "Incoming shared notes"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/shared-notes/incoming [get]
func (h *HandlerHTTP) getIncomingSharedNotes(c *gin.Context) {
	userID := c.MustGet("userID").(int64)

	notes, err := h.suc.GetIncomingSharedNotes(c, userID)
	if err != nil {
		log.Error("get all notes by notebook id: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrNoteNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.GetAllIncomingSharedNotesResponse{
		IncomingSharedNotes: notes,
		Count:               len(notes),
	})
}
