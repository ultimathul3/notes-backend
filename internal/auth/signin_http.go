package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

// @Summary	User sign in
// @Tags	Auth
// @Accept	json
// @Produce	json
// @Param	user body domain.GetUserDTO true "User JSON"
// @Success	200 {object} docs.SignInResponse "Data for user authorization"
// @Failure	400 {object} docs.MessageResponse "Error message"
// @Failure	500 {object} docs.MessageResponse "Server error message"
// @Router	/auth/sign-in [post]
func (h *HandlerHTTP) signIn(c *gin.Context) {
	var input domain.GetUserDTO
	if err := c.BindJSON(&input); err != nil {
		log.Error("GetUserIdDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
		log.Error("sign-in: user validation: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.uuc.Get(c, input)
	if err != nil {
		log.Errorf("get user '%s' (%s) by login and pass: %s", *input.Login, c.ClientIP(), err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.suc.GenerateTokens(user.ID)
	if err != nil {
		log.Error("generate tokens: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": domain.ErrInvalidOrExpiredAccessToken.Error(),
		})
		return
	}

	fingerprint := generateFingerprint(c)

	if h.suc.GetCountByUserID(c, user.ID) > h.suc.GetMaxUserSessionsCount()-1 {
		h.suc.DeleteAllByUserID(c, user.ID)
	}

	_, err = h.suc.Create(c, domain.CreateSessionDTO{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		Fingerprint:  fingerprint,
	})
	if err != nil {
		log.Error("create session: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	log.Infof("user '%s' (%s) has successfully logged in", *input.Login, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"id":            user.ID,
		"name":          user.Name,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
