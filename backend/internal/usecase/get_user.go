package usecase

import (
	"context"
)

// GetUser is the usecase for list camera annotation image.
type GetUser struct {
}

// GetUserInput is the input for GetUser.
type GetUserInput struct {
	UserID string `uri:"user_id" binding:"required"`
}

// GetUserOutput is the output structure for the camera annotation image usecase.
type GetUserOutput struct {
	UserID string
}

// Do executes the get camera annotation image usecase.
func (g GetUser) Do(
	_ context.Context,
	in GetUserInput,
) (GetUserOutput, error) {
	return GetUserOutput{UserID: in.UserID}, nil
}
