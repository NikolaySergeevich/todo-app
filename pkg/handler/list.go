package handler

import (
	"net/http"
	"strconv"

	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/gin-gonic/gin"
)

type getAllListsResponse struct {
	Lists []todoapp.TodoList `json:"Lists"`
}

type getAllItemsResponse struct {
	Items []todoapp.TodoItem `json:"Items"`
}

type getListByIdResponse struct {
	Data todoapp.TodoList `json:"Data"`
}

func (h *Handler) createlist(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todoapp.TodoList

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"listId": listId,
	})
}

func (h *Handler) getAlllist(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Lists: lists,
	})
}

func (h *Handler) getlistById(c *gin.Context){
	// этот метод ищет список по его id
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getListByIdResponse{
		Data: list,
	})
}

func (h *Handler) updateList(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	idList, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	binaryPayload, err := c.GetRawData()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "failed to get payload")
		return
	}

	var updatePayload todoapp.UpdateListPayload

	err = updatePayload.UnmarshalJSON(binaryPayload)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid payload")
	}

	err = h.services.Update(userId, idList, updatePayload)

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handler) deleteList(c *gin.Context){
	// этот метод удоляет список по его id
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}