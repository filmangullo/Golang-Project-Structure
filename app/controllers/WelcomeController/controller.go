package WelcomeController

import (
	"net/http"
	"your_project_name/app/constants"

	"github.com/gin-gonic/gin"
)

func WelcomeExecution(c *gin.Context) {
	var input WelcomeRequest

	// Parse and validate input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		constants.HandleErrorResponse(c, http.StatusBadRequest, "invalid json input", err.Error())
		return
	}

	// Parse and validate input Rule
	if err := ValidateRuleInput(input); err != nil {
		constants.HandleErrorResponse(c, http.StatusBadRequest, "invalid rule input", err.Error())
		return
	}

	// Call service logic
	results, err := WelcomeService(input)
	if err != nil {
		constants.HandleErrorResponse(c, http.StatusUnprocessableEntity, "Service execution failed", err.Error())
		return
	}

	// Return successful response
	response := constants.ResponseFormatter(http.StatusText(http.StatusOK), http.StatusOK, "success", results)
	c.JSON(http.StatusOK, response)
}
