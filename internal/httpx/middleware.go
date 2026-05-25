package httpx

import (
	"strings"

	"github.com/azmanabdlh/go-sample-api/internal/provider"
	"github.com/gin-gonic/gin"
)

func RequiredAuthentication(adapter provider.Adapter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value := strings.TrimSpace(
			ctx.GetHeader("Authorization"),
		)

		token := strings.TrimPrefix(
			value,
			"Bearer ",
		)

		if token == "" {
			RespondJSON(ctx, Response{
				Code:    401,
				Message: "Required Auth Token",
			})
			ctx.Abort()
			return
		}

		if err := adapter.ValidateToken(
			ctx.Request.Context(),
			token,
		); err != nil {
			RespondJSON(ctx, Response{
				Code:    401,
				Message: err.Error(),
			})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
