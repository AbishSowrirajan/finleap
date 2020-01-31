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

	r.HandleFunc("/cities/{id}", h.UpdateCity).Methods("PATCH")

	r.HandleFunc("/cities/{id}", h.DeleteCity).Methods("DELETE")

	r.HandleFunc("/temperatures", h.InsertTemperature).Methods("POST")

	r.HandleFunc("/forecasts/{id}", h.ForecastTemperature).Methods("GET")

	r.HandleFunc("/webhooks", h.CreateWebhooks).Methods("POST")

	r.HandleFunc("/webhooks/{id}", h.DeleteWebhooks).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
