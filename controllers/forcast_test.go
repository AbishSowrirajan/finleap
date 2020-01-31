package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/AbishSowrirajan/finleap/models"
	"github.com/stretchr/testify/assert"
)

type MockTempForcast struct {
	ID     string
	CityID string
	Max    float64
	Min    float64
	Sample string
}

type TestsTempForcast struct {
	Mockd       MockTempForcast
	Expected    int
	ExpectedMin float64
	ExpectedMax float64
	TestName    string
	URL         string
	Method      string
}

func TestForecastTemperature(t *testing.T) {

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
		{Expected: 201,
			TestName: "CreateCitySuccessfully",
			Method:   "POST",
			URL:      "/cities",
			Mockd: Mockdata{ID: "1",
				Name:      "Frankfut",
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
			Mockd: MockTemp{ID: "2",
				CityID:    "0",
				Max:       10.00,
				Min:       02.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
		{Expected: 201,
			TestName: "InsertTemperatureSuccessfully01",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "3",
				CityID:    "0",
				Max:       11.00,
				Min:       03.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
		{Expected: 201,
			TestName: "InsertTemperatureSuccessfully02",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "4",
				CityID:    "0",
				Max:       09.00,
				Min:       01.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
		{Expected: 201,
			TestName: "InsertTemperatureSuccessfully03",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "5",
				CityID:    "0",
				Max:       59.00,
				Min:       21.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
		{Expected: 201,
			TestName: "InsertTemperatureSuccessfully05",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "6",
				CityID:    "1",
				Max:       10.00,
				Min:       02.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
		{Expected: 201,
			TestName: "InsertTemperatureSuccessfully06",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "7",
				CityID:    "1",
				Max:       10.00,
				Min:       02.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
		{Expected: 201,
			TestName: "InsertTemperatureSuccessfully07",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "8",
				CityID:    "1",
				Max:       10.00,
				Min:       02.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
		{Expected: 201,
			TestName: "InsertTemperatureSuccessfully08",
			Method:   "POST",
			URL:      "/temperatures",
			Mockd: MockTemp{ID: "9",
				CityID:    "1",
				Max:       10.00,
				Min:       02.00,
				Timestamp: time.Now().UTC().Format("2006-01-02 03:04:05")}},
	}

	mockdataF := []TestsTempForcast{
		{Expected: 200,
			TestName:    "ForcastTempCitySuccessfullyCity1",
			Method:      "GET",
			URL:         "/forecasts/0",
			ExpectedMax: 22.25,
			ExpectedMin: 6.75,
			Mockd: MockTempForcast{ID: "0",
				CityID: "0",
				Max:    59.00,
				Min:    21.00,
				Sample: "4"}},
		{Expected: 200,
			TestName:    "ForcastTempCitySuccessfullyCity2",
			Method:      "GET",
			URL:         "/forecasts/1",
			ExpectedMax: 10.00,
			ExpectedMin: 02.00,
			Mockd: MockTempForcast{ID: "0",
				CityID: "0",
				Max:    59.00,
				Min:    21.00,
				Sample: "4"}},
		{Expected: 500,
			TestName:    "ForcastTempCityNotExist",
			Method:      "GET",
			URL:         "/forecasts/2",
			ExpectedMax: 0,
			ExpectedMin: 0,
			Mockd: MockTempForcast{ID: "0",
				CityID: "0",
				Max:    0,
				Min:    0,
				Sample: ""}},
	}

	for _, tt := range mockdata {

		t.Run(tt.TestName, InsertTempMock(tt, mockdb))

	}

	for _, tf := range mockdataF {

		t.Run(tf.TestName, ForecastTempMock(tf, mockdb))

	}

}

func ForecastTempMock(mockdata TestsTempForcast, mockdb MockDbdata) func(*testing.T) {

	return func(t *testing.T) {
		var result models.ForcastJSON

		request, _ := http.NewRequest(mockdata.Method, mockdata.URL, nil)
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		Router(mockdb).ServeHTTP(response, request)

		f, _ := ioutil.ReadAll(response.Body)

		_ = json.Unmarshal(f, &result)

		assert.Equal(t, mockdata.Expected, response.Code)

		minimum, _ := strconv.ParseFloat(result.Min, 64)
		maximum, _ := strconv.ParseFloat(result.Max, 64)

		assert.Equal(t, mockdata.ExpectedMin, minimum)
		assert.Equal(t, mockdata.ExpectedMax, maximum)
		assert.Equal(t, mockdata.Mockd.Sample, result.Sample)

	}
}
