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

	notebook := router.Group("/shared_notes").Use(tokenChecker)
	{
		notebook.POST("/", handler.create)
		notebook.DELETE("/:shared_note_id", handler.delete)
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
// @Router		/shared_notes [post]
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrNoteNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary		Deleting a shared note
// @Security	BearerToken
// @Tags		Shared note
// @Accept		json
// @Produce		json
// @Param		shared_note_id path int true "Shared note ID"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/shared_notes/{shared_note_id} [delete]
func (h *HandlerHTTP) delete(c *gin.Context) {
	sharedNoteID, err := strconv.ParseInt(c.Param("shared_note_id"), 10, 64)
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