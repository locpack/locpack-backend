package controller

import (
	"errors"

	"locpack-backend/pkg/adapter"
	"locpack-backend/pkg/types"
)

func getToken(ctx adapter.APIContext) (*types.Token, error) {
	rawToken, exists := ctx.Get("token")
	if !exists {
		return nil, nil
	}

	token, ok := rawToken.(types.Token)
	if !ok {
		return nil, errors.New("failed to parse token")
	}

	return &token, nil
}
