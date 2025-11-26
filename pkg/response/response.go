package response

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ToJson(c *gin.Context, status int, success bool, msg string, data interface{}) {
	c.JSON(status, ApiResponse{
		Success: success,
		Message: msg,
		Data:    data,
	})
}
