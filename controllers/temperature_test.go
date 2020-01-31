package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/AbishSowrirajan/finleap/models"
	"github.com/stretchr/testify/assert"
)

type MockTemp struct {
	ID        string
	CityID    string
	Max       float64
	Min       float64
	Timestamp string
}

type TestsTemp struct {
	Mockd    MockTemp
	Expected int
	TestName string
	URL      string
	Method   string
}

func TestInsertTemperature(t *testing.T) {

	mockdb := make(MockDbdata, 0)

	mockda := []tests{
		{Expected: 201,
			TestName: "CreateCitySuccessfully",
			Method:   "POST",
			URL:      "/cities",
			Mockd: Mockdata{ID: "0",
				Name:      "Berlin",
				Longitude: "10.00",
				Latitude:  "11.00"}},
	}

	for _, tt := range mockda {

		t.Run(tt.TestName, CreateCityMock(tt, mockdb))

	}

	mockdata := []TestsTemp{
		{Expected: 201,
			TestName: "InsertTemperatureSuccessfully",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "1",
				CityID:    "0",
				Max:       10.00,
				Min:       02.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
		{Expected: 201,
			TestName: "InsertTemperatureSuccessfully01",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "2",
				CityID:    "0",
				Max:       11.00,
				Min:       03.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
		{Expected: 201,
			TestName: "InsertTemperatureSuccessfully02",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "3",
				CityID:    "0",
				Max:       09.00,
				Min:       01.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
		{Expected: 400,
			TestName: "InsertTemperatureMinimumGreaterThanMaximum",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "",
				CityID:    "0",
				Max:       09.00,
				Min:       11.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
	}

	for _, tt := range mockdata {

		t.Run(tt.TestName, InsertTempMock(tt, mockdb))

	}

}

func InsertTempMock(mockdata TestsTemp, mockdb MockDbdata) func(*testing.T) {

	return func(t *testing.T) {
		var result models.TempJSON
		data := url.Values{}

		data.Set("city_id", mockdata.Mockd.CityID)
		data.Set("max", fmt.Sprintf("%2f", mockdata.Mockd.Max))
		data.Set("min", fmt.Sprintf("%2f", mockdata.Mockd.Min))

		request, _ := http.NewRequest(mockdata.Method, mockdata.URL, strings.NewReader(data.Encode()))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		Router(mockdb).ServeHTTP(response, request)

		f, _ := ioutil.ReadAll(response.Body)

		_ = json.Unmarshal(f, &result)

		assert.Equal(t, mockdata.Expected, response.Code)

		assert.Equal(t, mockdata.Mockd.ID, result.ID)

	}
}
