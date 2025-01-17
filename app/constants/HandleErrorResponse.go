package constants

import "github.com/gin-gonic/gin"

// handleErrorResponse sends a standardized error response
func HandleErrorResponse(c *gin.Context, statusCode int, message, errDetail string) {
	errorMessage := gin.H{"error": errDetail}
	response := ResponseFormatter(message, statusCode, "error", errorMessage)
	c.AbortWithStatusJSON(statusCode, response)
}
