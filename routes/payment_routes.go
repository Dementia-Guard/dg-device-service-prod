package routes

import (
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PaymentRoutes defines routes for payment-related operations
func PaymentRoutes(api *gin.RouterGroup) {
	paymentGroup := api.Group("/payment")
	{
		paymentGroup.GET("/", func(c *gin.Context) {
			utils.SuccessResponse(c, http.StatusOK, "Payment API Online", "SUCCESS", map[string]string{"message": "Payment Service Ready"})
		})
	}
}
