package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct{
	Timestamp string `json:"timestamp"`
	Code int `json:"code"`
	Message string `json:"message"`
	Status bool `json:"status"`
	Data interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context,code int, message string, status bool, data interface{}){
	c.JSON(code,Response{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Code: code,
		Message: message,
		Status: status,
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, code int,message string,status bool){
	c.JSON(code,Response{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Code:    code,
		Message:   message,
		Status:      status,
	})
}