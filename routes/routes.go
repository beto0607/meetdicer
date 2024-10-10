package routes

import "net/http"

func PrepareRouting() *http.ServeMux {

	apiRouter := prepareApiRouter()

	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", apiRouter))

	return router
}

func prepareApiRouter() *http.ServeMux {
	router := http.NewServeMux()
	return router
}
