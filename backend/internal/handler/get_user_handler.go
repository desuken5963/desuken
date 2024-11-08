package handler

import (
	"net/http"

	usecase "github.com/desuken/internal/usecase"
	"github.com/gin-gonic/gin"
)

// GetUserHandler is the handler for camera annotations.
type GetUserHandler struct {
	GetUserRunner GetUserRunner
}

func (h GetUserHandler) get(c *gin.Context) {
	ctx := c.Request.Context()
	var in usecase.GetUserInput

	if err := c.ShouldBindUri(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	result, err := h.GetUserRunner.Do(ctx, in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
