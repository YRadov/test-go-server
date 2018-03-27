package routers

import "github.com/gorilla/mux"

// InitRoutes registers all routes for the application.
func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the Bookmark entity
	router = SetBookmarkRoutes(router)
	// Call other router configurations
	return router
}
