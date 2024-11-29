package handler

import (
	"context"
	"net/http"

	// "github.com/desuken/internal/middleware"
	// "github.com/desuken/internal/routing"
	usecase "github.com/desuken/internal/usecase"
	"github.com/gin-gonic/gin"
)

// Options is options for NewHandler.
type Options struct {
	GetUserRunner  GetUserRunner
	GetOrderRunner GetOrderRunner
}

func NewHandler(opts Options) *gin.Engine {

	r := gin.New()

	getUserHandler := GetUserHandler{
		GetUserRunner: opts.GetUserRunner,
	}

	getOrderHandler := GetOrderHandler{
		GetOrderRunner: opts.GetOrderRunner,
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "health",
		})
	})

	r.GET("/user/:user_id", getUserHandler.get)

	r.GET("/order/:order_id", getOrderHandler.get)

	return r
}

// GetUserRunner is the interface of getUser usecase.
type GetUserRunner interface {
	Do(ctx context.Context, in usecase.GetUserInput) (usecase.GetUserOutput, error)
}

type GetOrderRunner interface {
	Do(ctx context.Context, in usecase.GetOrderInput) (usecase.GetOrderOutput, error)
}
