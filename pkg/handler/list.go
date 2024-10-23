package handler

import (
	"net/http"

	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/gin-gonic/gin"
)

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

}

func (h *Handler) getlistById(c *gin.Context){

}

func (h *Handler) updateList(c *gin.Context){

}

func (h *Handler) deleteList(c *gin.Context){

}