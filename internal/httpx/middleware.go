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
			ctx.JSON(401, gin.H{
				"message": "Required Auth Token",
			})

			ctx.Abort()
			return
		}

		if err := adapter.ValidateToken(
			ctx.Request.Context(),
			token,
		); err != nil {
			ctx.JSON(401, gin.H{
				"message": err.Error(),
			})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
