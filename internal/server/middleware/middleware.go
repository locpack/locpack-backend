package middleware

import (
	"net/http"
	"strings"

	"locpack-backend/internal/server/dto"
	"locpack-backend/pkg/adapter"
)

func AuthenticatedMiddleware(auth adapter.Auth) adapter.APIHandler {
	return func(ctx adapter.APIContext) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			errors := []dto.Error{{Message: "Some error", Code: "000"}}
			ctx.JSON(http.StatusUnauthorized, dto.ResponseWrapper{
				Meta:   dto.Meta{Success: false},
				Errors: errors,
			})
			return
		}

		accessToken := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := auth.DecodeToken(accessToken)
		if err != nil || !token.Valid {
			errors := []dto.Error{{Message: "Some error", Code: "000"}}
			ctx.JSON(http.StatusUnauthorized, dto.ResponseWrapper{
				Meta:   dto.Meta{Success: false},
				Errors: errors,
			})
			return
		}

		if err != nil {
			errors := []dto.Error{{Message: "Some error", Code: "000"}}
			ctx.JSON(http.StatusUnauthorized, dto.ResponseWrapper{
				Meta:   dto.Meta{Success: false},
				Errors: errors,
			})
			return
		}

		ctx.Set("myUserID", token.Username)

		ctx.Next()
	}
}

func NotAuthenticatedMiddleware() adapter.APIHandler {
	return func(ctx adapter.APIContext) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader != "" {
			errors := []dto.Error{{Message: "Some error", Code: "000"}}
			ctx.JSON(http.StatusUnauthorized, dto.ResponseWrapper{
				Meta:   dto.Meta{Success: false},
				Errors: errors,
			})
			return
		}

		ctx.Next()
	}
}
