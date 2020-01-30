package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AbishSowrirajan/finleap/models"
)

// Handler object to call routing object
type Handler struct {
	Db models.DbLayer
}

// NewHandler ...
func NewHandler() *Handler {

	handler := new(Handler)

	handler.Db = models.CityJSON{}

	return handler
}

// CreateCity ....
func (h *Handler) CreateCity(w http.ResponseWriter, r *http.Request) {

	var dbdata models.CityJSON

	r.ParseForm()

	dbdata.Name = r.PostFormValue("name")

	dbdata.Longitude = r.PostFormValue("longitude")

	dbdata.Latitude = r.PostFormValue("latitude")

	_ = h.Db.Create(dbdata)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("content-type", "application/json")
	result, _ := json.Marshal(dbdata)
	w.Write(result)

	return

}
