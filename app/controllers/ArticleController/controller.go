package ArticleController

import (
	"log"
	"net/http"
	"your_project_name/app/constants"
	"your_project_name/database"
	"your_project_name/database/tableArticle"
	"your_project_name/utils/StringsFunctions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
| Tujuan            | Gunakan                                |
| ----------------- | -------------------------------------- |
| Ambil body JSON   | `ShouldBindJSON()` (POST/PUT)          |
| Ambil query param | `ShouldBindQuery()` (GET)              |
| Ambil form-data   | `ShouldBind()` atau `ShouldBindForm()` |
*/

var db *gorm.DB
var tblArticle tableArticle.DatabaseTableArticle

func init() {
	var err error
	db, err = database.ConnectToDatabase()
	if err != nil {
		log.Printf("Error database connection :%v", err)
	}

	tblArticle = tableArticle.CallArticleRepository(db)
}

/**
 * CREATE
 */
func CreateExecution(c *gin.Context) {
	var input CreateRequest

	// Parse and validate input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		errorDetails := constants.FormatValidationError(err)
		response := constants.ResponseFormatter("invalid json input.", http.StatusBadRequest, "error", gin.H{"errors": errorDetails})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Parse and validate input Rule
	if err := ValidateCreateRequest(input); err != nil {
		response := constants.ResponseFormatter("invalid rule input.", http.StatusBadRequest, "error", gin.H{"errors": err})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Call service logic
	results, err := CreateService(input)
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

/**
 * LIST with paginate.
 */
func ListExecution(c *gin.Context) {
	var input ListRequest

	// Parse and validate input Query
	if err := c.ShouldBindQuery(&input); err != nil {
		errorDetails := constants.FormatValidationError(err)
		response := constants.ResponseFormatter("invalid query input.", http.StatusBadRequest, "error", gin.H{"errors": errorDetails})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Parse and validate input Rule
	if err := ValidateListRequest(input); err != nil {
		response := constants.ResponseFormatter("invalid rule input.", http.StatusBadRequest, "error", gin.H{"errors": err})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Call service logic
	results, err := ListArticleService(input)
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

	// Parse and validate slug path
	if !StringsFunctions.IsSlug(c.Param("slug")) {
		errorDetails := gin.H{
			"Slug": "must be lowercase, digits, and hyphens only",
		}
		response := constants.ResponseFormatter("invalid path format.", http.StatusBadRequest, "error", gin.H{"errors": errorDetails})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	} else {
		input.Slug = c.Param("slug")
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

func UpdateExecution(c *gin.Context) {
	var input UpdateRequest

	// Parse and validate input Query
	if err := c.ShouldBindJSON(&input); err != nil {
		errorDetails := constants.FormatValidationError(err)
		response := constants.ResponseFormatter("invalid query input.", http.StatusBadRequest, "error", gin.H{"errors": errorDetails})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Parse and validate slug path
	if !StringsFunctions.IsSlug(c.Param("id")) {
		errorDetails := gin.H{
			"ID": "must be a number and cannot be any other character",
		}
		response := constants.ResponseFormatter("invalid path format.", http.StatusBadRequest, "error", gin.H{"errors": errorDetails})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	} else {
		input.ID = c.Param("id")
	}

	// Parse and validate input Rule
	if err := ValidateUpdateRequest(input); err != nil {
		response := constants.ResponseFormatter("invalid rule input.", http.StatusBadRequest, "error", gin.H{"errors": err})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Call service logic
	results, err := UpdateArticleService(input)
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
