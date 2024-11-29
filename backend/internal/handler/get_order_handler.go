package handler

import (
	"net/http"

	usecase "github.com/desuken/internal/usecase"
	"github.com/gin-gonic/gin"
)

type GetOrderHandler struct {
	GetOrderRunner GetOrderRunner
}

func (h GetOrderHandler) get(c *gin.Context) {
	ctx := c.Request.Context()
	var in usecase.GetOrderInput

	if err := c.ShouldBindUri(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	result, err := h.GetOrderRunner.Do(ctx, in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
