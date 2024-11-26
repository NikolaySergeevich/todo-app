package handler

import (
	"fmt"
	"net/http"
	"strconv"

	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	binaryPayload, err := c.GetRawData()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get payload")
		return
	}

	var item todoapp.TodoItem
	err = item.UnmarshalJSON(binaryPayload)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invvalid payload")
		return
	}

	idItem, err := h.services.TodoItem.Create(userId, listId, item)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s: %s", "failed create item", err.Error()))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ItemId": idItem,
	})
}

func (h *Handler) getAllItem(c *gin.Context){

}

func (h *Handler) getItemById(c *gin.Context){

}

func (h *Handler) updateItem(c *gin.Context){

}

func (h *Handler) deleteItem(c *gin.Context){

}