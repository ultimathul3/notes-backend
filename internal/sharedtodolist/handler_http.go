package sharedtodolist

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type HandlerHTTP struct {
	suc domain.SharedTodoListUsecase
	uuc domain.UserUsecase
}

func NewHandlerHTTP(
	router *gin.Engine,
	suc domain.SharedTodoListUsecase,
	uuc domain.UserUsecase,
	tokenChecker gin.HandlerFunc,
) *HandlerHTTP {

	handler := &HandlerHTTP{
		suc: suc,
		uuc: uuc,
	}

	sharedNote := router.Group("/shared-todo-lists").Use(tokenChecker)
	{
		sharedNote.POST("/incoming", handler.create)
		sharedNote.GET("/", handler.getAllInfo)
		sharedNote.GET("/:shared-todo-list-id", handler.getDataByID)
		sharedNote.DELETE("/incoming/:shared-todo-list-id", handler.delete)
		sharedNote.POST("/incoming/:shared-todo-list-id", handler.accept)
		sharedNote.GET("/outgoing/:todo-list-id", handler.getOutgoingInfoByTodoListID)
	}

	return handler
}

// @Summary		Creating a shared todo list
// @Security	BearerToken
// @Tags		Shared todo list
// @Accept		json
// @Produce		json
// @Param		user body domain.CreateSharedTodoListDTO true "Shared todo list data"
// @Success		200 {object} docs.CreateSharedTodoListResponse "Shared todo list ID"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/shared-todo-lists/incoming [post]
func (h *HandlerHTTP) create(c *gin.Context) {
	var sharedList domain.CreateSharedTodoListDTO
	if err := c.BindJSON(&sharedList); err != nil {
		log.Error("CreateSharedTodoListDTO bind json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := sharedList.Validate(); err != nil {
		log.Error("CreateSharedTodoListDTO validate: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	whomID, err := h.uuc.GetUserIdByLogin(c, *sharedList.Login)
	if err != nil {
		log.Error("get user id by login: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrUserNotFound.Error()})
		return
	}

	userID := c.MustGet("userID").(int64)

	id, err := h.suc.Create(c, userID, whomID, *sharedList.TodoListID)
	if err != nil {
		log.Error("create shared todo list: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary		Deleting a shared todo list
// @Security	BearerToken
// @Tags		Shared todo list
// @Accept		json
// @Produce		json
// @Param		shared-todo-list-id path int true "Shared todo list ID"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/shared-todo-lists/incoming/{shared-todo-list-id} [delete]
func (h *HandlerHTTP) delete(c *gin.Context) {
	sharedTodoListID, err := strconv.ParseInt(c.Param("shared-todo-list-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid shared todo list id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	err = h.suc.Delete(c, sharedTodoListID, userID)
	if err != nil {
		log.Error("delete shared todo list: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrSharedTodoListNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @Summary		Getting a list of shared todo lists
// @Security	BearerToken
// @Tags		Shared todo list
// @Accept		json
// @Produce		json
// @Success		200 {array} domain.GetSharedTodoListsInfoResponse "Shared todo lists"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/shared-todo-lists [get]
func (h *HandlerHTTP) getAllInfo(c *gin.Context) {
	userID := c.MustGet("userID").(int64)

	lists, err := h.suc.GetAllInfo(c, userID)
	if err != nil {
		log.Error("get all shared todo lists info: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.GetSharedTodoListsInfoResponse{
		SharedTodoListsInfo: lists,
		Count:               len(lists),
	})
}

// @Summary		Accepting a shared todo list
// @Security	BearerToken
// @Tags		Shared todo list
// @Accept		json
// @Produce		json
// @Param		shared-todo-list-id path int true "Shared todo list ID"
// @Success		200 {object} docs.OkStatusResponse "OK status"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/shared-todo-lists/incoming/{shared-todo-list-id} [post]
func (h *HandlerHTTP) accept(c *gin.Context) {
	sharedTodoListID, err := strconv.ParseInt(c.Param("shared-todo-list-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid shared todo list id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	err = h.suc.Accept(c, sharedTodoListID, userID)
	if err != nil {
		log.Error("accept shared note: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrSharedTodoListNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @Summary		Getting a data of shared todo list
// @Security	BearerToken
// @Tags		Shared todo list
// @Accept		json
// @Produce		json
// @Param		shared-todo-list-id path int true "Shared todo list ID"
// @Success		200 {object} domain.SharedTodoListData "Data of shared todo list"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/shared-todo-lists/{shared-todo-list-id} [get]
func (h *HandlerHTTP) getDataByID(c *gin.Context) {
	sharedTodoListID, err := strconv.ParseInt(c.Param("shared-todo-list-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid shared todo list id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	data, err := h.suc.GetDataByID(c, sharedTodoListID, userID)
	if err != nil {
		log.Error("get data of shared todo list by id: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrSharedTodoListNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary		Getting an outgoing shared todo lists
// @Security	BearerToken
// @Tags		Shared todo list
// @Accept		json
// @Produce		json
// @Param		list-id path int true "Shared todo list ID"
// @Success		200 {array} domain.GetOutgoingSharedTodoListsInfoResponse "Outgoing shared todo lists"
// @Failure		400 {object} docs.MessageResponse "Error message"
// @Router		/shared-todo-lists/outgoing/{todo-list-id} [get]
func (h *HandlerHTTP) getOutgoingInfoByTodoListID(c *gin.Context) {
	listID, err := strconv.ParseInt(c.Param("todo-list-id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid todo list id param"})
		return
	}

	userID := c.MustGet("userID").(int64)

	lists, err := h.suc.GetOutgoingInfoByTodoListID(c, listID, userID)
	if err != nil {
		log.Error("get all shared todo lists info: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": domain.ErrSharedTodoListsNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.GetOutgoingSharedTodoListsInfoResponse{
		OutgoingSharedTodoListsInfo: lists,
		Count:                       len(lists),
	})
}
