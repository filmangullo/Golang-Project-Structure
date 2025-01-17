package middleware

import (
	"net/http"
	"your_project_name/app/constants"

	"github.com/gin-gonic/gin"
)

func AccessPermissionToken(c *gin.Context) {

	var tokenization = []map[string]string{
		{"access_name": "master-access", "access_token": "master-access-token"},
		{"access_name": "it-access", "access_token": "it-access-token"},
		// Add more valid tokens as needed.
	}

	token := c.GetHeader("AccessPermissionToken")

	tokenIsValid := false

	for _, t := range tokenization {
		if token == t["access_token"] {
			tokenIsValid = true
			break
		}
	}

	if !tokenIsValid {
		response := constants.ResponseFormatter("Still don't have permission", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

}
