package server

import (
	"net/http"
	"talentpro/base/router"

	"github.com/gorilla/handlers"
)

/**
*
* start the server
*
**/
func StartServer(port string) {
	headersOk := handlers.AllowedHeaders([]string{""})
	http.ListenAndServe(":"+port, handlers.CORS(headersOk)(router.HeadNodeRouter))
}
