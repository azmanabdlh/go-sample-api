package httpx

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func RespondJSON(c *gin.Context, response Response) {
	if response.Data == nil {
		response.Data = gin.H{}
	}

	c.JSON(response.Code, response)
}
