package routes

import (
	"api/controllers"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRoutes defines routes for user-related operations
func UserRoutes(api *gin.RouterGroup) {
	userGroup := api.Group("/user")
	{
		userGroup.GET("/", func(c *gin.Context) {
			utils.SuccessResponse(c, http.StatusOK, "User API Online", "SUCCESS", map[string]string{"message": "Hello From Go"})
		})
		userGroup.GET("/trigger-error", func(c *gin.Context) {
			panic("Something went wrong!")
		})
		userGroup.GET("/users", controllers.GetUsers)
	}
}
