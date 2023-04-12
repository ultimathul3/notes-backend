package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

func (h *HandlerHTTP) signIn(c *gin.Context) {
	var user domain.GetUserIDDTO
	if err := c.BindJSON(&user); err != nil {
		log.Error("GetUserIDDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := user.Validate(); err != nil {
		log.Errorf("sign-in: user validation: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := h.uuc.GetID(c, user)
	if err != nil {
		log.Errorf("get user '%s' (%s) by login and pass: %s", *user.Login, c.ClientIP(), err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.suc.GenerateTokens(id)
	if err != nil {
		log.Error("generate tokens: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": domain.ErrInvalidOrExpiredAccessToken.Error(),
		})
		return
	}

	fingerprint := generateFingerprint(c)

	if h.suc.GetCountByUserID(c, id) > h.suc.GetMaxUserSessionsCount()-1 {
		h.suc.DeleteAllByUserID(c, id)
	}

	_, err = h.suc.Create(c, domain.CreateSessionDTO{
		UserID:       id,
		RefreshToken: refreshToken,
		Fingerprint:  fingerprint,
	})
	if err != nil {
		log.Error("create session: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	log.Infof("user '%s' (%s) has successfully logged in", *user.Login, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"id":            id,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
