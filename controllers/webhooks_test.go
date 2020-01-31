package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/AbishSowrirajan/finleap/models"
	"github.com/stretchr/testify/assert"
)

type MockWebH struct {
	ID          string
	CityID      string
	CallBackURL string
}

type TestsWebH struct {
	Mockd    MockWebH
	Expected int
	TestName string
	URL      string
	Method   string
}

func TestCreateWebH(t *testing.T) {

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

	mockdbW := []TestsWebH{
		{Expected: 201,
			TestName: "CreateWebHooksSuccessfully",
			Method:   "POST",
			URL:      "/webhooks",
			Mockd: MockWebH{ID: "1",
				CityID:      "0",
				CallBackURL: "https://godoc.org/gopkg.in/go-playground/webhooks.v3"}},
		{Expected: 201,
			TestName: "CreateWebHooksSuccessfully",
			Method:   "POST",
			URL:      "/webhooks",
			Mockd: MockWebH{ID: "2",
				CityID:      "0",
				CallBackURL: "https://godoc.org/gopkg.in/go-playground/webhooks.v0"}},
		{Expected: 200,
			TestName: "DeleteWebHooksSuccessfully",
			Method:   "DELETE",
			URL:      "/webhooks/1",
			Mockd: MockWebH{ID: "1",
				CityID:      "0",
				CallBackURL: "https://godoc.org/gopkg.in/go-playground/webhooks.v3"}},
		{Expected: 400,
			TestName: "DeleteWebHooksNotEXist",
			Method:   "DELETE",
			URL:      "/webhooks/5",
			Mockd: MockWebH{ID: "",
				CityID:      "0",
				CallBackURL: "https://godoc.org/gopkg.in/go-playground/webhooks.v3"}},
	}

	for _, tt := range mockdbW {

		t.Run(tt.TestName, InsertWebHMock(tt, mockdb))

	}

}

func InsertWebHMock(mockdata TestsWebH, mockdb MockDbdata) func(*testing.T) {

	return func(t *testing.T) {
		var result models.WebHooksJSON
		data := url.Values{}

		data.Set("city_id", mockdata.Mockd.CityID)
		data.Set("callback_url", mockdata.Mockd.CallBackURL)

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
