package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AbishSowrirajan/finleap/models"
	"github.com/gorilla/mux"
)

// ForecastTemperature ....
func (h *Handler) ForecastTemperature(w http.ResponseWriter, r *http.Request) {

	var dbdata models.CityJSON
	var dbtemp models.ForcastJSON
	var output []byte

	r.ParseForm()

	vars := mux.Vars(r)

	dbdata.ID = vars["id"]

	result := h.Db.GetCity(dbdata)

	if _, ok := result.(models.CityJSON); ok {

		dbtemp.CityID = dbdata.ID

		result = h.Db.GetAvgTempByCity(dbtemp)

		if value, okay := result.(models.ForcastJSON); okay {

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
