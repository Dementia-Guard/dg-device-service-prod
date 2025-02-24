package routes

import (
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderRoutes defines routes for order-related operations
func OrderRoutes(api *gin.RouterGroup) {
	orderGroup := api.Group("/order")
	{
		orderGroup.GET("/", func(c *gin.Context) {
			utils.SuccessResponse(c, http.StatusOK, "Order API Online", "SUCCESS", map[string]string{"message": "Order Service Ready"})
		})
	}
}
