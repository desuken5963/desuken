package usecase

import (
	"context"
)

type GetOrder struct {
}

type GetOrderInput struct {
	OrderID string `uri:"order_id" binding:"required"`
}

type GetOrderOutput struct {
	OrderID string
}

func (g GetOrder) Do(
	_ context.Context,
	in GetOrderInput,
) (GetOrderOutput, error) {
	return GetOrderOutput{OrderID: in.OrderID}, nil
}
