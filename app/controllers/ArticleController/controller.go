package ArticleController

import (
	"net/http"
	"your_project_name/app/constants"

	"github.com/gin-gonic/gin"
)

/*
| Tujuan            | Gunakan                                |
| ----------------- | -------------------------------------- |
| Ambil body JSON   | `ShouldBindJSON()` (POST/PUT)          |
| Ambil query param | `ShouldBindQuery()` (GET)              |
| Ambil form-data   | `ShouldBind()` atau `ShouldBindForm()` |
*/

func PostExecution(c *gin.Context) {
	var input PostRequest

	// Parse and validate input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		errorDetails := constants.FormatValidationError(err)
		response := constants.ResponseFormatter("invalid json input.", http.StatusBadRequest, "error", gin.H{"errors": errorDetails})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Parse and validate input Rule
	if err := ValidatePostRequest(input); err != nil {
		response := constants.ResponseFormatter("invalid rule input.", http.StatusBadRequest, "error", gin.H{"errors": err})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Call service logic
	results, err := WelcomeService(input)
	if err != nil {
		errorDetails := constants.FormatValidationError(err)
		response := constants.ResponseFormatter("an error occurred while executing the service.", http.StatusBadRequest, "error", gin.H{"errors": errorDetails})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Return successful response
	response := constants.ResponseFormatter(http.StatusText(http.StatusOK), http.StatusOK, "success", results)
	c.JSON(http.StatusOK, response)
}

func GetExecution(c *gin.Context) {
	var input GetRequest

	// Parse and validate input Query
	if err := c.ShouldBindQuery(&input); err != nil {
		errorDetails := constants.FormatValidationError(err)
		response := constants.ResponseFormatter("invalid query input.", http.StatusBadRequest, "error", gin.H{"errors": errorDetails})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Parse and validate input Rule
	if err := ValidateGetRequest(input); err != nil {
		response := constants.ResponseFormatter("invalid rule input.", http.StatusBadRequest, "error", gin.H{"errors": err})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Call service logic
	results, err := GetArticleService(input)
	if err != nil {
		errorDetails := constants.FormatValidationError(err)
		response := constants.ResponseFormatter("an error occurred while executing the service.", http.StatusBadRequest, "error", gin.H{"errors": errorDetails})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Return successful response
	response := constants.ResponseFormatter(http.StatusText(http.StatusOK), http.StatusOK, "success", results)
	c.JSON(http.StatusOK, response)
}
