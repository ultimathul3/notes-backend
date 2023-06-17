package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

// @Summary		User logout
// @Security	BearerToken
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		user body domain.LogoutDTO true "User refresh token"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/auth/logout [post]
func (h *HandlerHTTP) logout(c *gin.Context) {
	var session domain.LogoutDTO
	if err := c.BindJSON(&session); err != nil {
		log.Error("LogoutDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)

	err := h.suc.Logout(c, userID, session)
	if err != nil {
		log.Error("user logout: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": domain.ErrInvalidOrExpiredRefreshToken.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
