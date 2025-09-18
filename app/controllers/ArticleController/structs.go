package ArticleController

import (
	articleResource "your_project_name/app/resources/articleResource"
	"your_project_name/utils/PaginateFunctions"
)

/*
|--------------------------------------------------------------------------
| Struct
|--------------------------------------------------------------------------
|
| This struct defines the structure of input data used across the application.
| It can include various fields required by an API endpoint, each of which
| may be validated using Gin's binding tags (e.g., `binding:"required"`).
|
| You can extend this struct to include any other necessary fields depending
| on the context in which it is used.
|
*/

/*
# -----------------------------------------------------------------------------
# CREATE
# -----------------------------------------------------------------------------
*/
type CreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Tags        string `json:"tags" binding:"required"`
	IsPublished *bool  `json:"isPublished" binding:"required"`
}

type CreateResponse struct {
	articleResource.ArticleArray
}

/*
# -----------------------------------------------------------------------------
# LIST
# -----------------------------------------------------------------------------
*/
type ListRequest struct {
	Page *int `form:"page"`
}

type ListResponse struct {
	Results  any                           `json:"results"`
	Page     int                           `json:"page"`
	PerPage  int                           `json:"per_page"`
	Total    int64                         `json:"total"`
	LastPage int                           `json:"last_page"`
	HasNext  bool                          `json:"has_next"`
	HasPrev  bool                          `json:"has_prev"`
	Labels   []PaginateFunctions.PageLabel `json:"labels"`
}

/*
# -----------------------------------------------------------------------------
# GET
# -----------------------------------------------------------------------------
*/
type GetRequest struct {
	Slug string `json:"slug"`
}

type GetResponse struct {
	Results []articleResource.ArticleArray `json:"results"`
}

/*
# -----------------------------------------------------------------------------
# UPDATE
# -----------------------------------------------------------------------------
*/
type UpdateRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Tags        string `json:"tags" binding:"required"`
	IsPublished *bool  `json:"isPublished" binding:"required"`
}

type UpdateResponse struct {
	Results articleResource.ArticleArray `json:"results"`
}

/*
# -----------------------------------------------------------------------------
# DELETE
# -----------------------------------------------------------------------------
*/
type DeleteRequest struct {
	ID string `json:"id"`
}

type DeleteResponse struct {
	Messages map[string]string `json:"messages"`
}
