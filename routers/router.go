package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	// summarize registration of handlers in mux router by role
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
