package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

func (h *HandlerHTTP) signIn(c *gin.Context) {
	var user *domain.GetUserIDDTO
	if err := c.BindJSON(&user); err != nil {
		log.Error("GetUserIDDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := h.uuc.GetID(c, user)
	if err != nil {
		log.Errorf("get user '%s' (%s) by login and pass: %s", *user.Login, c.ClientIP(), err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.jwt.GenerateTokens(id)
	if err != nil {
		log.Error("generate tokens: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": ErrInvalidOrExpiredAccessToken.Error()})
		return
	}

	fingerprint := fmt.Sprintf(
		"%s; %s; %s",
		c.ClientIP(),
		c.Request.Header["User-Agent"],
		c.Request.Header["Accept-Language"],
	)

	if h.suc.GetUserSessionsCount(c, id) > h.maxUserSessionsCount-1 {
		h.suc.DeleteAllUserSessions(c, id)
	}

	_, err = h.suc.Create(c, &domain.CreateSessionDTO{
		UserID:       id,
		RefreshToken: refreshToken,
		Fingerprint:  fingerprint,
		ExpiresIn:    time.Now().Add(h.refreshTokenTTL),
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