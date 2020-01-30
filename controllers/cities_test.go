package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/AbishSowrirajan/finleap/models"
	"github.com/gorilla/mux"
)

type MockDbdata struct {
	cities []string
}

func TestMain(t *testing.M) {

	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)

}

func Router() *mux.Router {

	r := mux.NewRouter()

	h := MockNewHandler()

	r.HandleFunc("/cities", h.CreateCity).Methods("POST")

	return r

}

func (data MockDbdata) Create(dbdata models.CityJSON) interface{} {

	fmt.Println("mockData")

	return nil
}

// MockNewHandler ...
func MockNewHandler() *Handler {

	handler := new(Handler)

	handler.Db = MockDbdata{}

	return handler
}

func TestCreateCity(t *testing.T) {

	data := url.Values{}

	data.Set("name", "Berlin")
	data.Set("longitude", "12,344")
	data.Set("latitude", "34.55")

	request, _ := http.NewRequest("POST", "/cities", strings.NewReader(data.Encode()))
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	f, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(f))

}
