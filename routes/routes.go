package routes

import (
	controllers_api "meetdicer/controllers/api"
	controllers_client "meetdicer/controllers/client"
	"net/http"
)

func PrepareRouting() *http.ServeMux {

	apiRouter := prepareApiRouter()
	clientRouter := prepareClientRouter()

	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", apiRouter))
	router.Handle("/client/", http.StripPrefix("/client", clientRouter))

	return router
}

func prepareApiRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /", controllers_api.GetIndex)
	return router
}

func prepareClientRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /", controllers_client.GetIndex)
	return router
}
