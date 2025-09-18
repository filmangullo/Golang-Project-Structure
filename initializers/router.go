package initializers

import (
	"your_project_name/app/controllers/ArticleController"
	"your_project_name/middleware"

	"github.com/gin-gonic/gin"
)

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application.
| so exact way of how its working is that you can
| - Api called -> Controller
| - Controller called -> Services
| - Services does the required things and also calls -> Database Connection like
|   fetch/update/delete etc dan gives the rresults.
|
*/
func Api() *gin.Engine {
	// Set Gin to release mode
	gin.SetMode(gin.DebugMode)

	// Create a Gin router instance
	router := gin.Default()

	// Apply CORS middleware
	router.Use(middleware.AllowCORS())
	// Set your trusted proxies here
	trustedProxies := []string{"192.168.0.1", "127.0.0.1", "45.77.173.85"} // Replace with actual proxy IPs
	if err := router.SetTrustedProxies(trustedProxies); err != nil {
		panic(err)
	}

	// Gin with the prefix versioning "V1"
	// Without access permissions
	api := router.Group("api/v1")
	// With access permissions
	// api := router.Group("api/v1", middleware.AccessPermissionToken)
	api.POST("articles", ArticleController.CreateExecution)       //create
	api.GET("articles", ArticleController.ListExecution)          //get
	api.GET("articles/:slug", ArticleController.GetExecution)     //find
	api.PUT("articles/:id", ArticleController.UpdateExecution)    //update
	api.DELETE("articles/:id", ArticleController.DeleteExecution) //delete

	// Define your API routes here

	return router
}
