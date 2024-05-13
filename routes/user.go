package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterModelAPIRoutes = func(r *mux.Router) {
	r.HandleFunc("/users/register", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	r.HandleFunc("/users/login", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
}
