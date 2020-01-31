package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AbishSowrirajan/finleap/models"
	"github.com/gorilla/mux"
)

// CreateWebhooks ....
func (h *Handler) CreateWebhooks(w http.ResponseWriter, r *http.Request) {

	var dbdata models.CityJSON
	var dbtemp models.WebHooksJSON
	//var customerror models.ModelError
	var output []byte

	r.ParseForm()

	dbdata.ID = r.PostFormValue("city_id")

	dbtemp.CallbackURL = r.PostFormValue("callback_url")

	result := h.Db.GetCity(dbdata)

	if _, ok := result.(models.CityJSON); ok {

		dbtemp.CityID = dbdata.ID

		result1 := h.Db.InsertCityWebH(dbtemp)

		if value, okay := result1.(models.WebHooksJSON); okay {

			result2 := h.Db.GetWebH(value)

			if value1, tr := result2.(models.WebHooksJSON); tr {

				output, _ = json.Marshal(value1)
				w.WriteHeader(http.StatusCreated)

			} else {

				val, _ := result2.(models.ModelError)

				output, _ = json.Marshal(val)

				if val.ErrTyp == "400" {

					w.WriteHeader(http.StatusBadRequest)
				} else {

					w.WriteHeader(http.StatusInternalServerError)
				}

			}

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

// DeleteWebhooks ....
func (h *Handler) DeleteWebhooks(w http.ResponseWriter, r *http.Request) {

	var dbtemp models.WebHooksJSON

	var output []byte

	vars := mux.Vars(r)

	dbtemp.ID = vars["id"]

	result := h.Db.DeleteWebH(dbtemp)

	if value, tr := result.(models.WebHooksJSON); tr {

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

	w.Header().Add("content-type", "application/json")

	w.Write(output)

	return

}
