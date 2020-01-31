package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/AbishSowrirajan/finleap/models"
)

// InsertTemperature ....
func (h *Handler) InsertTemperature(w http.ResponseWriter, r *http.Request) {

	var dbdata models.CityJSON
	var dbtemp models.TempJSON
	var customerror models.ModelError
	var output []byte

	r.ParseForm()

	dbdata.ID = r.PostFormValue("city_id")

	Max := r.PostFormValue("max")

	Min := r.PostFormValue("min")

	minimum, err := strconv.ParseFloat(Min, 3)

	if err != nil {

		customerror.Err = err.Error()
		customerror.ErrCode = "201"
		customerror.ErrTyp = "400"

		output, _ = json.Marshal(customerror)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("content-type", "application/json")

		w.Write(output)

		return

	}

	maximum, err := strconv.ParseFloat(Max, 3)

	if err != nil {

		customerror.Err = err.Error()
		customerror.ErrCode = "202"
		customerror.ErrTyp = "400"

		output, _ = json.Marshal(customerror)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("content-type", "application/json")

		w.Write(output)

		return

	}

	if minimum > maximum {
		customerror.Err = "Minimum temperature cannot be greater than Maximum temperaure"
		customerror.ErrCode = "203"
		customerror.ErrTyp = "400"

		output, _ = json.Marshal(customerror)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("content-type", "application/json")

		w.Write(output)

		return
	}

	dbtemp.Min = minimum
	dbtemp.Max = maximum

	dbtemp.CityID = dbdata.ID

	result := h.Db.GetCity(dbdata)

	if _, ok := result.(models.CityJSON); ok {

		now := time.Now().UTC().Format("2006-01-02 03:04:05")

		dbtemp.Timestamp = now

		result1 := h.Db.InsertCityTemp(dbtemp)

		if value, okay := result1.(models.TempJSON); okay {

			result2 := h.Db.GetTempByTimSt(value)

			if value1, tr := result2.(models.TempJSON); tr {

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
