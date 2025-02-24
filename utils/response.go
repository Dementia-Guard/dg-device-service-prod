package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct{
	Timestamp string `json:"timestamp"`
	Status int `json:"status"`
	Message string `json:"message"`
	Code string `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context,status int, message string, code string, data interface{}){
	c.JSON(status,Response{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Status: status,
		Message: message,
		Code: code,
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, status int,message string,code string){
	c.JSON(status,Response{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Status:    status,
		Message:   message,
		Code:      code,
	})
}