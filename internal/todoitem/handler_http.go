package todoitem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type HandlerHTTP struct {
	iuc domain.TodoItemUsecase
	luc domain.TodoListUsecase
}

func NewHandlerHTTP(router *gin.Engine, iuc domain.TodoItemUsecase, luc domain.TodoListUsecase, tokenChecker gin.HandlerFunc) *HandlerHTTP {
	handler := &HandlerHTTP{
		iuc: iuc,
		luc: luc,
	}

	todoItem := router.Group("/notebooks/:notebook-id/todo-lists/:todo-list-id/items").Use(tokenChecker)
	{
		todoItem.POST("/", handler.create)
		todoItem.GET("/", handler.getAllByListID)
		todoItem.PATCH("/:item-id", handler.patch)
		todoItem.DELETE("/:item-id", handler.delete)
	}

	return handler
}

// @Summary		Creating a todo item in todo list
// @Security	BearerToken
// @Tags		Todo item
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		todo-list-id path int true "Todo list ID"
// @Param		user body domain.CreateTodoItemDTO true "Item data"
// @Success		200 {object} docs.CreateTodoItemResponse "Todo list ID"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/notebooks/{notebook-id}/todo-lists/{todo-list-id}/items [post]
func (h *HandlerHTTP) create(c *gin.Context) {
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

	var item domain.CreateTodoItemDTO
	if err := c.BindJSON(&item); err != nil {
		log.Error("CreateTodoItemDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)

	id, err := h.iuc.Create(c, userID, notebookID, todoListID, item)
	if err != nil {
		log.Error("create todo list item: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrTodoListNotFound.Error()})
		return
	}

	if err = h.luc.RefreshUpdatedAt(c, todoListID, userID, notebookID); err != nil {
		log.Error("refresh todo list updated_at: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrTodoListNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary		Getting a list of items in todo list
// @Security	BearerToken
// @Tags		Todo item
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		todo-list-id path int true "Todo list ID"
// @Success		200 {array} docs.GetAllTodoItemsResponse "Todo items"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/notebooks/{notebook-id}/todo-lists/{todo-list-id}/items [get]
func (h *HandlerHTTP) getAllByListID(c *gin.Context) {
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

	items, err := h.iuc.GetAllByListID(c, userID, notebookID, todoListID)
	if err != nil {
		log.Error("get all todo items by list id: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrTodoListNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.GetAllTodoItemsResponse{
		TodoItems: items,
		Count:     len(items),
	})
}

// @Summary		Updating an item in todo list
// @Security	BearerToken
// @Tags		Todo item
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		todo-list-id path int true "Todo list ID"
// @Param		item-id path int true "Item ID"
// @Param		user body domain.PatchTodoItemDTO true "New todo item data"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/notebooks/{notebook-id}/todo-lists/{todo-list-id}/items/{item-id} [patch]
func (h *HandlerHTTP) patch(c *gin.Context) {
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

	todoItemID, err := strconv.ParseInt(c.Param("item-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid todo item id param"})
		return
	}

	var item domain.PatchTodoItemDTO
	if err := c.BindJSON(&item); err != nil {
		log.Error("UpdateTodoItemDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)

	err = h.iuc.Patch(c, todoItemID, userID, notebookID, todoListID, item)
	if err != nil {
		log.Error("patch todo item: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = h.luc.RefreshUpdatedAt(c, todoListID, userID, notebookID); err != nil {
		log.Error("refresh todo list updated_at: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrTodoListNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @Summary		Deleting an item from todo list
// @Security	BearerToken
// @Tags		Todo item
// @Accept		json
// @Produce		json
// @Param		notebook-id path int true "Notebook ID"
// @Param		todo-list-id path int true "Todo list ID"
// @Param		item-id path int true "Item ID"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/notebooks/{notebook-id}/todo-lists/{todo-list-id}/items/{item-id} [delete]
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

	todoItemID, err := strconv.ParseInt(c.Param("item-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid todo item id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	if err := h.iuc.Delete(c, todoItemID, userID, notebookID, todoListID); err != nil {
		log.Error("delete todo item: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrTodoListNotFound.Error()})
		return
	}

	if err := h.luc.RefreshUpdatedAt(c, todoListID, userID, notebookID); err != nil {
		log.Error("refresh todo list updated_at: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrTodoListNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
