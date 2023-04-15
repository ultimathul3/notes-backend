package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

// @Summary	Refreshing a user session
// @Tags	Auth
// @Accept	json
// @Produce	json
// @Param	user body docs.RefreshSessionDTO true "User session"
// @Success	200 {object} docs.RefreshResponse "New data for user authorization"
// @Failure	400 {object} docs.MessageResponse "Error message"
// @Router	/auth/refresh [post]
func (h *HandlerHTTP) refresh(c *gin.Context) {
	var input domain.RefreshSessionDTO
	if err := c.BindJSON(&input); err != nil {
		log.Error("RefreshSessionDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": domain.ErrInvalidOrExpiredRefreshToken.Error(),
		})
		return
	}

	if err := input.Validate(); err != nil {
		log.Error("refresh: input validation: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	input.Fingerprint = generateFingerprint(c)

	accessToken, refreshToken, err := h.suc.Refresh(c, input)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidFingerPrint) {
			log.Error("refresh tokens: token may have been stolen")
		} else {
			log.Error("refresh tokens: ", err)
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": domain.ErrInvalidOrExpiredRefreshToken.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
