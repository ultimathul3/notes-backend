package todolist

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type HandlerHTTP struct {
	tuc domain.TodoListUsecase
}

func NewHandlerHTTP(router *gin.Engine, tuc domain.TodoListUsecase, tokenChecker gin.HandlerFunc) *HandlerHTTP {
	handler := &HandlerHTTP{
		tuc: tuc,
	}

	todoList := router.Group("/notebooks/:notebook-id/todo-lists").Use(tokenChecker)
	{
		todoList.POST("/", handler.create)
		todoList.GET("/", handler.getAllByNotebookID)
		todoList.PUT("/:todo-list-id", handler.update)
		todoList.DELETE("/:todo-list-id", handler.delete)
	}

	return handler
}

// @Summary		Creating a todo list in a notebook
// @Security	BearerToken
// @Tags		Todo list
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		user body domain.CreateTodoListDTO true "Todo list data"
// @Success		200 {object} docs.CreateTodoListResponse "Todo list ID"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id}/todo-lists/ [post]
func (h *HandlerHTTP) create(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	var list domain.CreateTodoListDTO
	if err := c.BindJSON(&list); err != nil {
		log.Error("CreateTodoListDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)

	id, err := h.tuc.Create(c, userID, notebookID, list)
	if err != nil {
		log.Error("create note: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrNotebookNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary		Getting a list of user todo lists in a notebook
// @Security	BearerToken
// @Tags		Todo list
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Success		200 {array} docs.GetAllTodoListsResponse "Todo lists"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id}/todo-lists/ [get]
func (h *HandlerHTTP) getAllByNotebookID(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	lists, err := h.tuc.GetAllByNotebookID(c, userID, notebookID)
	if err != nil {
		log.Error("get all todo lists by notebook id: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrNotebookNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.GetAllTodoListsResponse{
		TodoLists: lists,
		Count:     len(lists),
	})
}

// @Summary		Updating a todo list in a notebook
// @Security	BearerToken
// @Tags		Todo list
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		todo-list-id path int true "Todo list ID"
// @Param		user body domain.UpdateTodoListDTO true "New todo list data"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id}/todo-lists/{todo-list-id} [put]
func (h *HandlerHTTP) update(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	todoListID, err := strconv.ParseInt(c.Param("todo-list-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid todo list id param"})
		return
	}

	var list domain.UpdateTodoListDTO
	if err := c.BindJSON(&list); err != nil {
		log.Error("UpdateTodoListDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)

	err = h.tuc.Update(c, todoListID, userID, notebookID, list)
	if err != nil {
		log.Error("update todo list: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @Summary		Deleting a todo list from a notebook
// @Security	BearerToken
// @Tags		Todo list
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		todo-list-id path int true "Todo list ID"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/api/notebooks/{notebook-id}/todo-lists/{todo-list-id} [delete]
func (h *HandlerHTTP) delete(c *gin.Context) {
	notebookID, err := strconv.ParseInt(c.Param("notebook-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid notebook id param"})
		return
	}

	todoListID, err := strconv.ParseInt(c.Param("todo-list-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid todo list id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	err = h.tuc.Delete(c, todoListID, userID, notebookID)
	if err != nil {
		log.Error("update note: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrTodoListNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
