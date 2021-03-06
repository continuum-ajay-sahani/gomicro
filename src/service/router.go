package service

import (
	"net/http"
	"practice/gomicro/src/config"
	"practice/gomicro/src/lib"

	"github.com/gorilla/mux"
)

// NewRouter creates a router for URL-to-service mapping
func NewRouter() *mux.Router {
	config := config.Configuration
	libRouter := lib.GetReqRouter()
	router := libRouter.StrictSlash(true)

	api := router.PathPrefix(config.URLPrefix).Subrouter()
	//apiV1 := api.PathPrefix(config.APIVersion).Subrouter()

	api.Handle("/version", lib.WrapMethodHandler(http.MethodGet, handlerAppVersion))
	api.Handle("/health", lib.WrapMethodHandler(http.MethodGet, handlerAppHealth))

	// Consumer APIs
	// Use for retrieving features on different levels for specific partner

	return router
}
