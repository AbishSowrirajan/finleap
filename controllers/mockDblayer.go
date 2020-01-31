package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/AbishSowrirajan/finleap/models"
	"github.com/gorilla/mux"
)

type MockDbdata map[string]interface{}

func Router(mockdb MockDbdata) *mux.Router {

	r := mux.NewRouter()

	h := MockNewHandler(mockdb)

	r.HandleFunc("/cities", h.CreateCity).Methods("POST")
	r.HandleFunc("/cities/{id}", h.UpdateCity).Methods("PATCH")
	r.HandleFunc("/cities/{id}", h.DeleteCity).Methods("DELETE")
	r.HandleFunc("/temperatures", h.InsertTemperature).Methods("POST")
	r.HandleFunc("/forecasts/{id}", h.ForecastTemperature).Methods("GET")
	r.HandleFunc("/webhooks", h.CreateWebhooks).Methods("POST")

	r.HandleFunc("/webhooks/{id}", h.DeleteWebhooks).Methods("DELETE")

	return r

}

func (citydb MockDbdata) Create(dbdata models.CityJSON) interface{} {

	var customerror models.ModelError

	result := citydb.GetCityByName(dbdata)

	_, ok := result.(models.CityJSON)

	if ok {

		customerror.Err = "City already exist"
		customerror.ErrCode = "101"
		customerror.ErrTyp = "400"

		return customerror

	}

	length := len(citydb)

	dbdata.ID = strconv.Itoa(length)

	citydb[dbdata.ID] = dbdata

	return dbdata

}

func (citydb MockDbdata) GetCity(dbdata models.CityJSON) interface{} {

	var customerror models.ModelError

	if val, ok := citydb[dbdata.ID]; ok {

		return val
	}

	customerror.Err = "City doesnt exist"
	customerror.ErrCode = "105"
	customerror.ErrTyp = "400"

	return customerror
}

func (citydb MockDbdata) UpdateCity(dbdata models.CityJSON) interface{} {

	if _, ok := citydb[dbdata.ID]; ok {

		citydb[dbdata.ID] = dbdata
	}

	return dbdata
}

func (citydb MockDbdata) DeleteCity(dbdata models.CityJSON) interface{} {

	if _, ok := citydb[dbdata.ID]; ok {

		delete(citydb, dbdata.ID)
	}

	return dbdata
}

func (citydb MockDbdata) InsertCityTemp(dbdata models.TempJSON) interface{} {

	if _, ok := citydb[dbdata.ID]; !ok {

		length := len(citydb)

		dbdata.ID = strconv.Itoa(length)

		citydb[dbdata.ID] = dbdata
	}

	time.Sleep(1 * time.Second)

	return dbdata
}

func (citydb MockDbdata) GetCityByName(dbdata models.CityJSON) interface{} {

	var customerror models.ModelError

	for _, value := range citydb {

		val, _ := value.(models.CityJSON)

		if val.Name == dbdata.Name {

			return value
		}

	}

	customerror.Err = "City doesnt exist"
	customerror.ErrCode = "107"
	customerror.ErrTyp = "400"

	return customerror

}

func (citydb MockDbdata) GetTempByTimSt(dbdata models.TempJSON) interface{} {

	var customerror models.ModelError

	for _, value := range citydb {

		val, _ := value.(models.TempJSON)

		if val.Timestamp == dbdata.Timestamp {

			return val
		}

	}

	customerror.Err = "Temperature belongs to Timestamp doesnt exist"
	customerror.ErrCode = "107"
	customerror.ErrTyp = "400"

	return customerror

}

func (citydb MockDbdata) GetAvgTempByCity(dbdata models.ForcastJSON) interface{} {

	min := 0.00
	max := 0.00
	sample := 0
	for _, val := range citydb {

		if value, okay := val.(models.TempJSON); okay && value.CityID == dbdata.CityID {

			sample++
			min += value.Min
			max += value.Max

		}
	}

	dbdata.Max = fmt.Sprintf("%2f", max/float64(sample))
	dbdata.Min = fmt.Sprintf("%2f", min/float64(sample))
	dbdata.Sample = strconv.Itoa(sample)

	return dbdata
}

func (citydb MockDbdata) InsertCityWebH(dbdata models.WebHooksJSON) interface{} {

	if _, ok := citydb[dbdata.ID]; !ok {

		length := len(citydb)

		dbdata.ID = strconv.Itoa(length)

		citydb[dbdata.ID] = dbdata
	}

	return dbdata
}

func (citydb MockDbdata) GetWebH(dbdata models.WebHooksJSON) interface{} {

	var customerror models.ModelError

	for _, val := range citydb {

		if value, okay := val.(models.WebHooksJSON); okay && value.CityID == dbdata.CityID && value.CallbackURL == dbdata.CallbackURL {

			dbdata.ID = value.ID

			return dbdata
		}
	}

	customerror.Err = "Callbask URL belong to this city doesnt exist"
	customerror.ErrCode = "107"
	customerror.ErrTyp = "400"

	return customerror
}

func (citydb MockDbdata) DeleteWebH(dbdata models.WebHooksJSON) interface{} {

	var customerror models.ModelError

	if value, ok := citydb[dbdata.ID]; ok {

		dbdata, _ = value.(models.WebHooksJSON)

		delete(citydb, dbdata.ID)

		return dbdata

	}

	customerror.Err = "Callbask URL belong to this city doesnt exist"
	customerror.ErrCode = "107"
	customerror.ErrTyp = "400"

	return customerror
}

// MockNewHandler ...
func MockNewHandler(mockdb MockDbdata) *Handler {

	handler := new(Handler)

	handler.Db = mockdb

	return handler
}
