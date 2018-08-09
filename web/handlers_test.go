package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNotFoundError(t *testing.T) {
	r := httptest.NewRequest("GET", "/password/3/1", nil)
	resp := httptest.NewRecorder()
	newRouter := NewRouter()
	newRouter.NotFoundHandler = http.HandlerFunc(NotFoundError)
	newRouter.ServeHTTP(resp, r)

	if resp.Code != 404 {
		t.Errorf("Handler didn't return the proper status Code. Returned status Code %d", resp.Code)
	}

	if resp.Header().Get("Content-Type") != "application/json; charset=UTF-8" {
		t.Errorf("Handler didn't return the proper Content Type. "+
			"returned Content Type %s", resp.Header().Get("Content-Type"))
	}

	bodyResponse := strings.TrimRight(resp.Body.String(), "\n")
	expectedBodyResponse := `{"error":"Not Found","message":"The defined route '/password/3/1' has not been found","type":"404"}`
	if bodyResponse != expectedBodyResponse {
		t.Errorf("Handler didn't return the proper Body %s", bodyResponse)
	}
}

func TestValidationError(t *testing.T) {
	r := httptest.NewRequest("GET", "/password/3/1/1/false_route", nil)
	resp := httptest.NewRecorder()
	newRouter := NewRouter()
	newRouter.ServeHTTP(resp, r)

	if resp.Code != 422 {
		t.Errorf("Handler didn't return the proper status Code. Returned status Code %d", resp.Code)
	}

	if resp.Header().Get("Content-Type") != "application/json; charset=UTF-8" {
		t.Errorf("Handler didn't return the proper Content Type. "+
			"returned Content Type %s", resp.Header().Get("Content-Type"))
	}

	bodyResponse := strings.TrimRight(resp.Body.String(), "\n")
	expectedBodyResponse := `{"error":"Validation Error","message":"The defined params are not valid or not allowed","type":"422"}`
	if bodyResponse != expectedBodyResponse {
		t.Errorf("Handler didn't return the proper Body %s", bodyResponse)
	}
}

func TestGetPassword(t *testing.T) {
	r := httptest.NewRequest("GET", "/password/3/1/1/1", nil)
	resp := httptest.NewRecorder()
	newRouter := NewRouter()
	newRouter.ServeHTTP(resp, r)

	if resp.Code != 200 {
		t.Errorf("Handler didn't return the proper status Code. Returned status Code %d", resp.Code)
	}

	if resp.Header().Get("Content-Type") != "application/json; charset=UTF-8" {
		t.Errorf("Handler didn't return the proper Content Type. "+
			"returned Content Type %s", resp.Header().Get("Content-Type"))
	}
}
