package routers

import (
	"net/http"

	"github.com/AbishSowrirajan/finleap/controllers"
	"github.com/gorilla/mux"
)

// Run the routing Handlers ....
func Run() {

	r := mux.NewRouter()

	h := controllers.NewHandler()

	r.HandleFunc("/cities", h.CreateCity).Methods("POST")

	http.ListenAndServe(":8080", r)

}
