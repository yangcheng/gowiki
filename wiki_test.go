package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPageNameValidatd(t *testing.T) {

	req, err := http.NewRequest("GET", "/edit/normal123", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(makeHandler(editHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestPageNameInvalid(t *testing.T) {

	req, err := http.NewRequest("GET", "/edit/normal123&*", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(makeHandler(editHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

}
