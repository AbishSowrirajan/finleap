package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AbishSowrirajan/finleap/models"
	"github.com/gorilla/mux"
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
	var output []byte
	var customerror models.ModelError

	r.ParseForm()

	dbdata.Name = r.PostFormValue("name")

	dbdata.Longitude = r.PostFormValue("longitude")

	dbdata.Latitude = r.PostFormValue("latitude")

	if len(dbdata.Name) == 0 || len(dbdata.Longitude) == 0 || len(dbdata.Latitude) == 0 {

		customerror.Err = "Input field cannot be Empty"
		customerror.ErrCode = "200"
		customerror.ErrTyp = "400"

		output, _ = json.Marshal(customerror)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("content-type", "application/json")

		w.Write(output)

		return
	}

	result := h.Db.Create(dbdata)

	if _, ok := result.(models.CityJSON); ok {

		result1 := h.Db.GetCityByName(dbdata)

		if val, okay := result1.(models.CityJSON); okay {

			output, _ = json.Marshal(val)
			w.WriteHeader(http.StatusCreated)

		} else {
			val, _ := result1.(models.ModelError)

			output, _ = json.Marshal(val)

			if val.ErrTyp == "400" {

				w.WriteHeader(http.StatusBadRequest)
			} else {

				w.WriteHeader(http.StatusInternalServerError)
			}

		}

	} else {

		val, _ := result.(models.ModelError)

		output, _ = json.Marshal(val)

		if val.ErrTyp == "400" {

			w.WriteHeader(http.StatusBadRequest)
		} else {

			w.WriteHeader(http.StatusInternalServerError)
		}

	}

	w.Header().Add("content-type", "application/json")

	w.Write(output)

	return

}

// UpdateCity ....
func (h *Handler) UpdateCity(w http.ResponseWriter, r *http.Request) {

	var dbdata models.CityJSON
	var output []byte
	var customerror models.ModelError

	r.ParseForm()

	vars := mux.Vars(r)

	dbdata.ID = vars["id"]

	result := h.Db.GetCity(dbdata)

	if val, ok := result.(models.CityJSON); ok {

		dbdata.Name = r.PostFormValue("name")

		dbdata.Longitude = r.PostFormValue("longitude")

		dbdata.Latitude = r.PostFormValue("latitude")

		if len(dbdata.Name) == 0 || len(dbdata.Longitude) == 0 || len(dbdata.Latitude) == 0 {

			customerror.Err = "Input field cannot be Empty"
			customerror.ErrCode = "201"
			customerror.ErrTyp = "400"

			output, _ = json.Marshal(customerror)

			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("content-type", "application/json")

			w.Write(output)

			return
		}

		if val == dbdata {

			customerror.Err = "New data is same as old"
			customerror.ErrCode = "202"
			customerror.ErrTyp = "400"

			output, _ = json.Marshal(customerror)

			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("content-type", "application/json")

			w.Write(output)

			return

		}

		result = h.Db.UpdateCity(dbdata)

		if value, okay := result.(models.CityJSON); okay {

			output, _ = json.Marshal(value)
			w.WriteHeader(http.StatusCreated)

		} else {

			val, _ := result.(models.ModelError)

			output, _ = json.Marshal(val)

			if val.ErrTyp == "400" {

				w.WriteHeader(http.StatusBadRequest)
			} else {

				w.WriteHeader(http.StatusInternalServerError)
			}

		}

	} else {

		val, _ := result.(models.ModelError)

		output, _ = json.Marshal(val)

		if val.ErrTyp == "400" {

			w.WriteHeader(http.StatusBadRequest)
		} else {

			w.WriteHeader(http.StatusInternalServerError)
		}

	}

	w.Header().Add("content-type", "application/json")

	w.Write(output)

	return

}

// DeleteCity ....
func (h *Handler) DeleteCity(w http.ResponseWriter, r *http.Request) {

	var dbdata models.CityJSON
	var output []byte

	r.ParseForm()

	vars := mux.Vars(r)

	dbdata.ID = vars["id"]

	result := h.Db.GetCity(dbdata)

	if val, ok := result.(models.CityJSON); ok {

		result = h.Db.DeleteCity(val)

		if value, okay := result.(models.CityJSON); okay {

			output, _ = json.Marshal(value)
			w.WriteHeader(http.StatusOK)

		} else {

			val, _ := result.(models.ModelError)

			output, _ = json.Marshal(val)

			if val.ErrTyp == "400" {

				w.WriteHeader(http.StatusBadRequest)
			} else {

				w.WriteHeader(http.StatusInternalServerError)
			}

		}

	} else {

		val, _ := result.(models.ModelError)

		output, _ = json.Marshal(val)

		if val.ErrTyp == "400" {

			w.WriteHeader(http.StatusBadRequest)
		} else {

			w.WriteHeader(http.StatusInternalServerError)
		}

	}

	w.Header().Add("content-type", "application/json")

	w.Write(output)

	return

}
