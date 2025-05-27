package ArticleController

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
# POST
# -----------------------------------------------------------------------------
*/
type PostRequest struct {
	// Add fields here if needed. For example:
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Content  string `json:"content" binding:"required"`
	PostView int    `json:"postView" binding:"required"`
}

type PostResponse struct {
}

/*
# -----------------------------------------------------------------------------
# GET
# -----------------------------------------------------------------------------
*/
type GetRequest struct {
	// Add fields here if needed. For example:
	Slug string `form:"slug" binding:"required"`
}

type GetResponse struct {
}
