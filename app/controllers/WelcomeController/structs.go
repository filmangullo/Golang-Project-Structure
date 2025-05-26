package WelcomeController

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

type WelcomeRequest struct {
	// Add fields here if needed. For example:
	Name string `json:"name" binding:"required"`
}

type WelcomeResponse struct {
}
