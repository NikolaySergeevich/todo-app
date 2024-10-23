package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx = "userId"
)

func(h *Handler) userIdentity(c *gin.Context) {
	header :=  c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParsed := strings.Split(header, " ")
	if len(headerParsed) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userID, err := h.services.Authorization.ParseToken(headerParsed[1])
	if err != nil {
		newErrorResponse(c,http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)
}

// из контекста с аутентификацией деуствующей получается id пользователя типа interface{} и приводиться к типу int
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid tipy")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}