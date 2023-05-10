package search

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type HandlerHTTP struct {
	suc domain.SearchUsecase
}

func NewHandlerHTTP(router *gin.Engine, suc domain.SearchUsecase, tokenChecker gin.HandlerFunc) *HandlerHTTP {
	handler := &HandlerHTTP{
		suc: suc,
	}

	search := router.Group("/search").Use(tokenChecker)
	{
		search.GET("/", handler.search)
	}

	return handler
}

// @Summary		Search in notes, todo lists, shared notes and shared todo lists
// @Security	BearerToken
// @Tags		Search
// @Accept		json
// @Produce		json
// @Param		title query string true "Search by title"
// @Param		notes query boolean false "Search by notes"
// @Param		todo-lists query boolean false "Search by todo lists"
// @Success		200 {array} domain.SearchResult "Search result"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/search [get]
func (h *HandlerHTTP) search(c *gin.Context) {
	search := domain.Search{}

	userID := c.MustGet("userID").(int64)

	title := c.Query("title")
	search.Title = title

	byNotes := c.Query("notes")
	if byNotes == "" {
		search.ByNotes = true
	} else {
		b, err := strconv.ParseBool(byNotes)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notes param"})
			return
		}
		search.ByNotes = b
	}

	byTodoLists := c.Query("todo-lists")
	if byTodoLists == "" {
		search.ByTodoLists = true
	} else {
		b, err := strconv.ParseBool(byTodoLists)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid todo-lists param"})
			return
		}
		search.ByTodoLists = b
	}

	result, err := h.suc.GetAll(c, userID, search)
	if err != nil {
		log.Error("search error: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
