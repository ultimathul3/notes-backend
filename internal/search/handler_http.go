package search

import (
	"net/http"
	"strconv"
	"time"

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
// @Param		title query string false "Search by title"
// @Param		notes query boolean false "Search by notes"
// @Param		todo-lists query boolean false "Search by todo lists"
// @Param		shared-notes query boolean false "Search by shared notes"
// @Param		shared-todo-lists query boolean false "Search by shared todo lists"
// @Param		created-from query number false "Created from (timestamp)"
// @Param		created-to query number false "Created to (timestamp)"
// @Param		updated-from query number false "Updated from (timestamp)"
// @Param		updated-to query number false "Updated to (timestamp)"
// @Success		200 {array} domain.SearchResult "Search result"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/search [get]
func (h *HandlerHTTP) search(c *gin.Context) {
	search := domain.Search{}

	userID := c.MustGet("userID").(int64)

	search.Title = c.Query("title")

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

	bySharedNotes := c.Query("shared-notes")
	if bySharedNotes == "" {
		search.BySharedNotes = true
	} else {
		b, err := strconv.ParseBool(bySharedNotes)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid shared-notes param"})
			return
		}
		search.BySharedNotes = b
	}

	bySharedTodoLists := c.Query("shared-todo-lists")
	if bySharedTodoLists == "" {
		search.BySharedTodoLists = true
	} else {
		b, err := strconv.ParseBool(bySharedTodoLists)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid shared-todo-lists param"})
			return
		}
		search.BySharedTodoLists = b
	}

	createdFrom := c.Query("created-from")
	if createdFrom != "" {
		timestamp, err := strconv.ParseInt(createdFrom, 10, 64)
		if err != nil || timestamp < 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid created-from param"})
			return
		}
		search.CreatedFrom = time.Unix(timestamp, 0)
	}

	createdTo := c.Query("created-to")
	if createdTo != "" {
		timestamp, err := strconv.ParseInt(createdTo, 10, 64)
		if err != nil || timestamp < 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid created-to param"})
			return
		}
		search.CreatedTo = time.Unix(timestamp, 0)
	}

	updatedFrom := c.Query("updated-from")
	if updatedFrom != "" {
		timestamp, err := strconv.ParseInt(updatedFrom, 10, 64)
		if err != nil || timestamp < 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid updated-from param"})
			return
		}
		search.UpdatedFrom = time.Unix(timestamp, 0)
	}

	updatedTo := c.Query("updated-to")
	if updatedTo != "" {
		timestamp, err := strconv.ParseInt(updatedTo, 10, 64)
		if err != nil || timestamp < 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid updated-to param"})
			return
		}
		search.UpdatedTo = time.Unix(timestamp, 0)
	}

	result, err := h.suc.GetAll(c, userID, search)
	if err != nil {
		log.Error("search error: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
