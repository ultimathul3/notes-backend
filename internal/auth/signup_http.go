package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

func (h *HandlerHTTP) signUp(c *gin.Context) {
	var user *domain.CreateUserDTO
	if err := c.BindJSON(&user); err != nil {
		log.Error("CreateUserDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := h.uuc.Create(c, user)
	if err != nil {
		log.Error("create user: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Infof("user '%s' (%s) has been created", *user.Login, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"id": id})
}
