package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	//Initialize variables
	test_data := api_req{}
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	//Send Request
	w := httptest.NewRecorder()
	test_data.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	//Validate Response
	if res.StatusCode != 200 {
		t.Errorf("Error Code Received. ")
	}
}

func TestRequestCount(t *testing.T) {
	//Initialize variables
	test_data := api_req{}
	os.Truncate("request_log.csv", 0)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	w1 := httptest.NewRecorder()

	//Send Req
	test_data.ServeHTTP(w, req)
	test_data.ServeHTTP(w, req)
	test_data.ServeHTTP(w, req)
	test_data.ServeHTTP(w1, req)
	res := w1.Result()
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//Validate Response
	if string(body) != "No of request in last 1 min is: 4" {
		t.Errorf("The count doesn't match. The actual response is: " + string(body))
	}

}
