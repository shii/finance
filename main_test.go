package main

import (
	"encoding/json"
	"fmt"
	"github.com/shii/finance/model"
	"github.com/shii/finance/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var route *http.ServeMux
var res *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	c := model.SetupCache()
	route = routes.SetupRoutes(c)
	res = httptest.NewRecorder()
	code := m.Run()
	os.Exit(code)
}

func TestGetStatusCode200(t *testing.T) {

	m := routes.TransferRequest{AccountFrom: 1234, AccountTo: 7890, Amount: 100}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.NewReader(string(b))
	req, err := http.NewRequest(http.MethodPost, "/v1/transfer", body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	route.ServeHTTP(res, req)
	fmt.Printf("%v", res.Code)
	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestGetStatusCode400(t *testing.T) {

	m := routes.TransferRequest{AccountFrom: 1234, AccountTo: 7890, Amount: 1200}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.NewReader(string(b))
	req, err := http.NewRequest(http.MethodPost, "/v1/transfer", body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	route.ServeHTTP(res, req)
	fmt.Printf("%v", res.Code)
	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}
