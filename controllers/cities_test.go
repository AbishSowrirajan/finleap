package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/AbishSowrirajan/finleap/models"
	"github.com/stretchr/testify/assert"
)

type Mockdata struct {
	ID        string
	Name      string
	Longitude string
	Latitude  string
}

type tests struct {
	Mockd    Mockdata
	Expected int
	TestName string
	URL      string
	Method   string
}

func TestMain(m *testing.M) {

	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)

}

func TestCreateCity(t *testing.T) {

	mockdb := make(MockDbdata, 0)

	mockdata := []tests{
		{Expected: 201,
			TestName: "CreateCitySuccessfully",
			Method:   "POST",
			URL:      "/cities",
			Mockd: Mockdata{ID: "0",
				Name:      "Berlin",
				Longitude: "10.00",
				Latitude:  "11.00"}},
		{Expected: 400,
			TestName: "CreateDuplicateCity",
			Method:   "POST",
			URL:      "/cities",
			Mockd: Mockdata{ID: "",
				Name:      "Berlin",
				Longitude: "10.00",
				Latitude:  "11.00"}},
		{Expected: 400,
			TestName: "CreateEmptyInputField",
			Method:   "POST",
			URL:      "/cities",
			Mockd: Mockdata{ID: "",
				Name:      "Frankfut",
				Longitude: "",
				Latitude:  "11.00"}},
		{Expected: 405,
			TestName: "CreateUnknownMethod",
			Method:   "GET",
			URL:      "/cities",
			Mockd: Mockdata{ID: "",
				Name:      "Frankfut",
				Longitude: "10.00",
				Latitude:  "11.00"}},
		{Expected: 201,
			TestName: "UpdateCitySuccessfully",
			Method:   "PATCH",
			URL:      "/cities/0",
			Mockd: Mockdata{ID: "0",
				Name:      "Frankfut",
				Longitude: "10.00",
				Latitude:  "11.00"}},
		{Expected: 400,
			TestName: "UpdateCityNotExist",
			Method:   "PATCH",
			URL:      "/cities/1",
			Mockd: Mockdata{ID: "",
				Name:      "Frankfut",
				Longitude: "10.00",
				Latitude:  "11.00"}},
		{Expected: 400,
			TestName: "UpdateCitySameAsOld",
			Method:   "PATCH",
			URL:      "/cities/0",
			Mockd: Mockdata{ID: "",
				Name:      "Frankfut",
				Longitude: "10.00",
				Latitude:  "11.00"}},
		{Expected: 400,
			TestName: "UpdateCityWithEmptyField",
			Method:   "PATCH",
			URL:      "/cities/0",
			Mockd: Mockdata{ID: "",
				Name:      "Frankfut",
				Longitude: "",
				Latitude:  "11.00"}},
		{Expected: 200,
			TestName: "DeleteCitySuccessfully",
			Method:   "DELETE",
			URL:      "/cities/0",
			Mockd: Mockdata{ID: "0",
				Name:      "Frankfut",
				Longitude: "10.00",
				Latitude:  "11.00"}},
		{Expected: 400,
			TestName: "DeleteCityNotExit",
			Method:   "DELETE",
			URL:      "/cities/0",
			Mockd: Mockdata{ID: "",
				Name:      "Frankfut",
				Longitude: "10.00",
				Latitude:  "11.00"}},
	}

	for _, tt := range mockdata {

		t.Run(tt.TestName, CreateCityMock(tt, mockdb))

	}

}

func CreateCityMock(mockdata tests, mockdb MockDbdata) func(*testing.T) {

	return func(t *testing.T) {
		var result models.CityJSON
		data := url.Values{}

		data.Set("name", mockdata.Mockd.Name)
		data.Set("longitude", mockdata.Mockd.Longitude)
		data.Set("latitude", mockdata.Mockd.Latitude)

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
